import React, { useEffect, useState } from 'react';
import { createRoot } from 'react-dom/client';
import Markdown from 'react-markdown';
import remarkGfm from 'remark-gfm';
import rehypeSlug from 'rehype-slug';
import rehypeHighlight from 'rehype-highlight';
import { documents, contentUrl, catPreview } from './content';
import { loadStars } from './stars';
import { repository } from './project';
import './style.css';
import ThemeToggle from './ThemeToggle';
import DocPagination from './DocPagination';
import gopher from '../../assets/gopher.png';
import wisprFlowPreview from '../../assets/wispr-flow.png';
import teamsPreview from '../../assets/teams.jpg';
import dinoPreview from '../../assets/dino.png';

const demos = [
  {
    title: 'Wispr Flow', poster: wisprFlowPreview,
    description: 'HTML and CSS screens rendered by a layout engine written in ownframe.',
    video: `${import.meta.env.BASE_URL}demos/preview.mp4`,
    post: 'https://x.com/chinmay_sawant_/status/2106788230154871126',
  },
  {
    title: 'Desktop cat overlay', poster: catPreview,
    description: 'A transparent, click-through cat that reports what is happening inside opencode.',
    video: `${import.meta.env.BASE_URL}demos/desktop-cat.mp4`,
    post: 'https://x.com/chinmay_sawant_/status/2106829998409789929',
  },
  {
    title: 'Teams demo', poster: teamsPreview,
    description: 'A native chat and collaboration app built with Go, HTML, and CSS using ownframe.',
    video: `${import.meta.env.BASE_URL}demos/teams.mp4`,
    post: 'https://x.com/chinmay_sawant_/status/2107548496194895993/video/1',
  },
  {
    title: 'Dino Run', poster: dinoPreview,
    description: 'An endless runner built with HTML, CSS, and Go using ownframe.',
    example: `${repository}/tree/master/examples/dino`,
  },
];
const quickStart = `go get github.com/chinmay-sawant/ownframe

# Run an example from the checkout
go run ./examples/login`;

function App() {
  const [hash, setHash] = useState(location.hash);
  const [stars, setStars] = useState(null);
  useEffect(() => {
    let mounted = true;
    loadStars().then((value) => { if (mounted) setStars(value); });
    const update = () => setHash(location.hash);
    window.addEventListener('hashchange', update);
    return () => { mounted = false; window.removeEventListener('hashchange', update); };
  }, []);
  const [, slug = 'readme', anchor] = hash.split('/');
  const isDocs = hash.startsWith('#docs');
  const document = documents.find((entry) => entry.slug === slug);
  useEffect(() => {
    if (isDocs && anchor) {
      documentElement(anchor)?.scrollIntoView();
    } else if (isDocs || hash === '' || hash === '#home') window.scrollTo(0, 0);
    else documentElement(hash.slice(1))?.scrollIntoView();
  }, [hash, isDocs, anchor]);

  return <>
    <a className="skip" href="#main">Skip to content</a>
    <header>
      <a className="brand" href="#home">ownframe</a>
      <nav aria-label="Main navigation">
        <a href="#demos">Demos</a>
        <a href="#docs/readme" aria-current={isDocs ? 'page' : undefined}>Documentation</a>
      </nav>
      <ThemeToggle />
      <a className="github" href={repository} title="View ownframe on GitHub">
        <span className="github-star" aria-hidden="true">⭐</span>
        GitHub <span aria-live="polite">{stars === null ? '' : stars.toLocaleString() + ' stars'}</span>
      </a>
    </header>
    <main id="main" tabIndex="-1">
      {isDocs ? <div className="docs-layout">
        <aside>
          <h2>Documentation</h2>
          <nav aria-label="Documentation topics">{documents.map((entry) =>
            <a key={entry.slug} href={`#docs/${entry.slug}`} aria-current={slug === entry.slug ? 'page' : undefined}>{entry.title}</a>
          )}</nav>
        </aside>
        <article className="markdown">
          {document ? <><Markdown remarkPlugins={[remarkGfm]} rehypePlugins={[rehypeSlug, rehypeHighlight]}
            urlTransform={(url) => contentUrl(url, document)}>{document.text}</Markdown>
            <DocPagination documents={documents} slug={slug} /></>
            : <><h1>Document not found</h1><a href="#docs/readme">Open the documentation index</a></>}
        </article>
      </div> : <>
        <section className="intro">
          <img className="mascot" src={gopher} width="180" height="180" alt="A cheerful blue Gopher waving hello" />
          <h1>HTML screens. Go logic.</h1>
          <p>ownframe was formerly known as go-gpui.</p>
          <p>Write your screen in HTML and CSS, handle its data and actions in Go, and run it on desktop, in the browser through WebAssembly, or on a phone.</p>
          <p>ownframe uses <a href="https://github.com/chinmay-sawant/gowkhtmltopdf">gowkhtmltopdf</a> for layout and Ebiten for the window. It ships no Chromium, WebKit, or JavaScript runtime.</p>
          <div className="links"><a href="#start">Get started</a><a href="#docs/readme">Read the documentation</a></div>
        </section>
        <section id="demos">
          <h2>Demos</h2>
          <div className="demos">{demos.map((demo) => <article key={demo.title}>
            <h3>{demo.title}</h3><p>{demo.description}</p>
            {demo.video ? <video controls playsInline preload="metadata" poster={demo.poster} aria-label={demo.title}>
              <source src={demo.video} type="video/mp4" />
              Your browser does not support this video. Use the original post below.
            </video> : <img className="demo-image" src={demo.poster} loading="lazy" alt="Dino Run in a native desktop window" />}
            <a href={demo.post || demo.example}>{demo.post ? 'Watch the original video on X' : 'View the Dino example on GitHub'}</a>
          </article>)}</div>
        </section>
        <section id="start">
          <h2>Get started</h2><pre><code>{quickStart}</code></pre>
          <p>The login demo accepts <code>secret</code> for both fields. <a href="#docs/examples-readme">Browse the examples</a> or read about <a href="#docs/window">opening a window</a>.</p>
        </section>
        <section>
          <h2>Documentation</h2>
          <p>The guides cover <a href="#docs/binding">data binding</a>, <a href="#docs/theming">theming</a>, <a href="#docs/platforms">platforms</a>, and <a href="#docs/packaging">packaging</a>.</p>
          <p>There is no live DOM or JavaScript engine, and the library does not support video, canvas, or WebGL. See the <a href="#docs/features">feature list and limits</a>.</p>
        </section>
      </>}
    </main>
    <footer><a href={repository}>Source on GitHub</a><a href="#docs/readme">Documentation</a></footer>
  </>;
}

function documentElement(id) { return window.document.getElementById(id); }
createRoot(document.getElementById('root')).render(<React.StrictMode><App /></React.StrictMode>);
