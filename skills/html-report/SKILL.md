---
name: html-report
description: >
  Turn a topic, an analysis, or an agent's own output into one self-contained
  HTML page a person can understand: controlled-language prose, inline SVG
  diagrams, interactive detail, no build step, no network. Use when the user
  says "html report", "turn this into a web page", "explain this in HTML",
  "build an explainer", "make this readable", or runs /html-report. Escalates
  to an explainer video on request (reference/explainer-video.md). Not for
  PDF output or golden fixtures (gowkhtmltopdf tooling), and not for
  verifying a diff (diff-verify).
---

# HTML report

A chat reply is the worst format for an explanation. The reader scrolls, the
diagram is 400 lines up, and nothing can be moved. One HTML file puts the
prose, the diagrams, and the controls in one place. It opens from `file://`,
needs no server, and gets thrown away when it stops being useful.

The ladder of output formats, ranked by what the reader gets:

1. Prose. Fast to write, hard to reread.
2. Diagrams. Good for shape and flow, silent on detail.
3. An HTML page. Prose, diagrams, tables, and interaction in one file.
4. An explainer video. Best for things that move over time.

Each rung costs more to build and saves the reader more effort. The default
stop is rung 3, which already folds in rungs 1 and 2. Rung 4 is a branch:
`reference/explainer-video.md`.

## Invocation

- `/html-report <topic>`: explain a topic. Ground it in sources first.
- `/html-report this`: turn the current conversation or analysis into a page.
- `/html-report <path>`: report on a file, a diff, or command output.
- `/html-report --video <topic>`: build an explainer video instead; read
  `reference/explainer-video.md`.
- `--out <path>`: write the file to a chosen path.

Default output path: `plans/reports/html/<slug>-<date>.html`.

## 1. Fix the question

Write down three things before building anything:

- One sentence: what the reader must understand.
- One reader: the actual person who will open the file. "A new contributor"
  beats "anyone".
- The sources: files with line numbers, or the exact places in the
  conversation.

When the material lives in the repo, read the code, tests, or plans first,
and collect each `file:line` at the moment you read it. Never explain from
memory.

Done when: one sentence, one reader, and every factual claim has a named
source.

## 2. Outline by the reader's questions

Order sections the way the reader will ask them, not the way the source
presents them. Start with why the thing exists, then what it does, then how
it works, then the edges. For each section pick its form:

| Reader needs | Form |
|---|---|
| to compare three or more things | table |
| to see flow, branching, or nesting | diagram |
| to move a number and watch the result | control |
| a definition or a caveat | prose |
| the exact code or strings | formatted code block |
| a side path not every reader takes | collapsible `<details>` |

Done when: every section answers a question the one reader would ask, and no
section is a source inventory.

## 3. Write in controlled language

Target about 80 percent of ASD-STE100, the controlled-language spec written
for aerospace maintenance. Most of its value sits in a short list:

- One idea per sentence. Twelve words is normal, twenty is the ceiling.
- Active voice, present tense. "The window draws the frame", not "the frame
  is drawn".
- One term per concept, used every time. Name a thing once, then never call
  it by another word.
- Define a term in the sentence where it first appears.
- No metaphor as explanation. "under the hood" and "handles that" name no
  mechanism.
- One instruction per sentence, started by a verb.
- Fact first, then reason, then exception.
- Numbers over adjectives. "3.1 ms" beats "fast".
- No em dashes.

Done when: no sentence needs a second read, and every term is plain or
defined on the spot.

## 4. Draw the hard parts

Replace any "how it flows" paragraph with an inline SVG diagram.

- One idea per diagram. Split the idea instead of adding to the drawing.
- Label every part with the name the code uses, so the diagram doubles as a
  map.
- Give it one caption sentence that states the idea.
- Keep the page's text and line colors. No stock art, no shadows.

Done when: the diagram still tells a true story with the surrounding prose
deleted.

## 5. Build one self-contained file

- One `.html` file with inline `<style>` and `<script>`. No CDN, no fetched
  font, no network request at any point. It must open from `file://`.
- Semantic markup: `main`, `section`, `figure`, `figcaption`, `table`,
  `details`.
- Light and dark themes through `prefers-color-scheme`.
- Body text 16 to 18 px at a 60 to 70 character measure.
- The first render reads with JavaScript off. Use JS for controls only.
- Interactive where it teaches: `<details>` for side paths, tabs for
  alternatives, one slider or button for a parameter. Hover and focus
  highlight the related part.
- A tl;dr of at most five lines at the top, linking to each section.
- Print styles that hide controls and keep diagrams.
- Under about 200 KB. Bigger means something entered that is not the
  explanation.
- Taste reference: `rendering-flow.html` at the repo root. Warm paper,
  cards, one accent per concept, no nav chrome.

Skeleton to copy:

```html
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>TOPIC</title>
<style>
  :root { --bg:#f6f3ec; --card:#fff; --ink:#1c1915; --muted:#6f675a;
          --line:#ddd5c6; --accent:#1a56db; }
  @media (prefers-color-scheme: dark) {
    :root { --bg:#15130f; --card:#1e1b16; --ink:#ece7dd; --muted:#a49b8b;
            --line:#3a352c; --accent:#7aa2ff; }
  }
  body { margin:0; background:var(--bg); color:var(--ink);
         font:17px/1.55 system-ui, sans-serif; }
  main { max-width:70ch; margin:0 auto; padding:2.5rem 1.25rem 5rem; }
  .card { background:var(--card); border:1px solid var(--line);
          border-radius:12px; padding:1rem 1.25rem; }
  svg { max-width:100%; height:auto; }
  @media print { .control { display:none; } }
</style>
</head>
<body>
<main>
  <h1>TOPIC</h1>
  <p class="card">tl;dr in at most five lines.</p>
  <section id="why"><h2>Why it exists</h2><p>...</p></section>
</main>
</body>
</html>
```

Done when: opening the file from `file://` gives zero console errors and
zero network requests.

## 6. Read it like a reader

- Screenshot the first screen and the full page, then read both images:

```bash
google-chrome --headless --screenshot=/tmp/opencode/first.png \
  --window-size=1440,900 "file://$PWD/plans/reports/html/<name>.html"
google-chrome --headless --screenshot=/tmp/opencode/full.png \
  --window-size=1440,5600 "file://$PWD/plans/reports/html/<name>.html"
```

- Confirm the first screen states the question and the tl;dr.
- Switch to dark mode and read the first screen again.
- Turn JavaScript off and confirm the page still reads.
- Prove nothing loads from the network:
  `grep -nE '<link|<script src|@import|url\(http' <file>` must print
  nothing. A citation URL lives in text, not in a tag.
- Check that every `file:line` citation still exists.

Done when: both screenshots were read, the console is clean, and every
citation exists.

## 7. Deliver

- Write the file, then reply with its path, the one-line question, and the
  source count.
- Do not paste the page into chat.
- Let it go. The page is disposable; the sources outlive it.

## Definition of done

- One file, opens offline from `file://`, no network requests, no console
  errors.
- The one reader can answer the one-line question from the page alone.
- Every claim about the repo cites a `file:line` that exists.
- The prose follows the controlled-language list.
- Each diagram carries real names and one caption.
- Light, dark, print, and no-JS all read.
- The unslop pass ran over the page copy.

## Related skills

- `skills/feynman/SKILLS.md`: when the understanding itself is not there
  yet.
- `skills/diff-verify/SKILL.md`: verify claims against commit history before
  publishing them.
