import { test } from 'node:test';
import assert from 'node:assert/strict';
import { createStarLoader, cacheKey } from './stars.js';

const hour = 3600000;
function fixture() {
  const values = new Map();
  return {
    storage: { getItem: (key) => values.get(key) ?? null, setItem: (key, value) => values.set(key, value) },
    now: () => 100000000,
    fetcher: async () => new Response(JSON.stringify({ stargazers_count: 42 })),
  };
}

test('shares concurrent requests and reuses persisted counts across page loads for six hours', async () => {
  const deps = fixture();
  let calls = 0;
  deps.fetcher = async () => { calls++; return new Response('{"stargazers_count":42}'); };
  const load = createStarLoader(deps);
  assert.deepEqual(await Promise.all([load(), load(), load()]), [42, 42, 42]);
  assert.equal(await createStarLoader(deps)(), 42);
  assert.equal(calls, 1);
  deps.now = () => 100000000 + 6 * hour;
  assert.equal(await createStarLoader(deps)(), 42);
  assert.equal(calls, 2);
});

test('keeps stale count on rate limit and honors the server backoff', async () => {
  const deps = fixture();
  deps.storage.setItem(cacheKey, JSON.stringify({ count: 25, nextRequest: 0 }));
  let calls = 0;
  deps.fetcher = async () => {
    calls++;
    return new Response('', { status: 429, headers: { 'retry-after': '7200' } });
  };
  assert.equal(await createStarLoader(deps)(), 25);
  deps.now = () => 100000000 + hour;
  assert.equal(await createStarLoader(deps)(), 25);
  assert.equal(calls, 1);
  assert.equal(JSON.parse(deps.storage.getItem(cacheKey)).nextRequest, 100000000 + 2 * hour);
});

test('honors a longer rate-limit reset and HTTP-date Retry-After', async () => {
  for (const headers of [
    { 'x-ratelimit-reset': String((100000000 + 4 * hour) / 1000) },
    { 'retry-after': new Date(100000000 + 4 * hour).toUTCString() },
  ]) {
    const deps = fixture();
    deps.fetcher = async () => new Response('', { status: 403, headers });
    assert.equal(await createStarLoader(deps)(), null);
    assert.equal(JSON.parse(deps.storage.getItem(cacheKey)).nextRequest, 100000000 + 4 * hour);
  }
});

test('network and malformed response failures are cached for an hour without inventing a count', async () => {
  for (const fetcher of [
    async () => { throw new Error('offline'); },
    async () => new Response('{"stargazers_count":"42"}'),
    async () => new Response('not JSON'),
  ]) {
    const deps = fixture();
    let calls = 0;
    deps.fetcher = async () => { calls++; return fetcher(); };
    assert.equal(await createStarLoader(deps)(), null);
    assert.equal(await createStarLoader(deps)(), null);
    assert.equal(calls, 1);
  }
});

test('blocked or corrupted storage still permits a count and uses memory caching', async () => {
  for (const storage of [
    { getItem: () => '{broken', setItem: () => { throw new Error('blocked'); } },
    { getItem: () => { throw new Error('blocked'); }, setItem: () => { throw new Error('blocked'); } },
  ]) {
    let calls = 0;
    const load = createStarLoader({ ...fixture(), storage,
      fetcher: async () => { calls++; return new Response('{"stargazers_count":0}'); },
    });
    assert.equal(await load(), 0);
    assert.equal(await load(), 0);
    assert.equal(calls, 1);
  }
});

test('tab lock rechecks shared storage after another tab refreshes', async () => {
  const deps = fixture();
  let queue = Promise.resolve();
  deps.lock = (_name, task) => {
    const result = queue.then(task);
    queue = result.catch(() => {});
    return result;
  };
  let calls = 0;
  deps.fetcher = async () => { calls++; return new Response('{"stargazers_count":42}'); };
  const firstTab = createStarLoader(deps);
  const secondTab = createStarLoader(deps);
  assert.deepEqual(await Promise.all([firstTab(), secondTab()]), [42, 42]);
  assert.equal(calls, 1);
});
