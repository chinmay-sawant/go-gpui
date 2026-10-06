import React from 'react';

export default function DocPagination({ documents, slug }) {
  const index = documents.findIndex((entry) => entry.slug === slug);
  if (index < 0) return null;
  const previous = documents[index - 1];
  const next = documents[index + 1];

  return <nav className="doc-pagination" aria-label="Documentation pages">
    {previous && <a href={`#docs/${previous.slug}`} rel="prev">
      <span>← Previous</span><strong>{previous.title}</strong>
    </a>}
    {next && <a href={`#docs/${next.slug}`} rel="next">
      <span>Next →</span><strong>{next.title}</strong>
    </a>}
  </nav>;
}
