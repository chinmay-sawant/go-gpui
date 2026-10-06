import { starsEndpoint } from './project.js';
export const cacheKey = 'ownframe:github-stars:v1';
const hour = 60 * 60 * 1000;

export function createStarLoader({ storage, fetcher, now = Date.now, lock } = {}) {
  let memory = { count: null, nextRequest: 0 };
  let pending;

  function read() {
    try {
      const value = JSON.parse(storage?.getItem(cacheKey) ?? storage?.getItem('go-gpui:github-stars:v1'));
      if (value && (value.count === null || (Number.isSafeInteger(value.count) && value.count >= 0))
        && Number.isFinite(value.nextRequest) && value.nextRequest >= 0) memory = value;
    } catch { /* Storage can be unavailable or contain an invalid entry. */ }
    return memory;
  }

  function save(value) {
    memory = value;
    try { storage?.setItem(cacheKey, JSON.stringify(value)); } catch { /* Keep the memory cache. */ }
    return value.count;
  }

  async function refresh() {
    const cached = read();
    if (cached.nextRequest > now()) return cached.count;
    // Persist a cooldown before the request, including if this tab closes mid-fetch.
    save({ count: cached.count, nextRequest: now() + hour });
    let nextRequest = now() + hour;
    try {
      const response = await fetcher(starsEndpoint, {
        headers: { Accept: 'application/vnd.github+json' },
        signal: AbortSignal.timeout(10000),
      });
      if (!response.ok) {
        const retry = response.headers.get('retry-after');
        const reset = response.headers.get('x-ratelimit-reset');
        const retryTime = retry === null ? 0
          : /^\d+$/.test(retry) ? now() + Number(retry) * 1000 : Date.parse(retry);
        const resetTime = reset === null ? 0 : Number(reset) * 1000;
        nextRequest = Math.max(nextRequest, retryTime || 0, resetTime || 0);
        throw new Error('GitHub count unavailable');
      }
      const data = await response.json();
      if (!Number.isSafeInteger(data.stargazers_count) || data.stargazers_count < 0) {
        throw new Error('Invalid GitHub count');
      }
      return save({ count: data.stargazers_count, nextRequest: now() + 6 * hour });
    } catch {
      return save({ count: cached.count, nextRequest });
    }
  }

  return function loadStars() {
    if (!pending) {
      pending = (lock ? lock('ownframe:github-stars', refresh) : refresh())
        .catch(() => read().count)
        .finally(() => { pending = undefined; });
    }
    return pending;
  };
}

let storage;
try { storage = globalThis.localStorage; } catch { /* Browser privacy settings. */ }
export const loadStars = createStarLoader({
  storage,
  fetcher: (...args) => fetch(...args),
  lock: globalThis.navigator?.locks
    ? (name, callback) => navigator.locks.request(name, callback) : undefined,
});
