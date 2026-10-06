---
name: explain-output
description: >
  Render a topic, an analysis, a diff, or a model answer in the output format
  that makes it easiest to understand: controlled-language prose, an inline
  diagram, one self-contained HTML page, or a narrated explainer video. Use
  when the user asks to explain or visualize something, says an answer is hard
  to read, asks for a report page or a diagram, or asks for a video or
  3b1b-style explainer.
---

# Explain output

A model's answer only helps when the reader can take it in. The output format
decides how much work that takes. A chat reply is the worst format for an
explanation. The reader scrolls, the diagram scrolls away, and every fact
carries the same weight.

Build the strongest format the question justifies, then let the artifact go.
The artifact is disposable; the sources outlive it.

## 1. Fix the question

Write three things before building anything:

- One sentence: what the reader must understand.
- One reader: the actual person who will open the artifact. "A new
  contributor" beats "anyone".
- The sources: files with line numbers, or the exact places in the
  conversation.

When the material lives in the repo, read the code, tests, or plans first, and
collect each `file:line` at the moment you read it. Never explain from memory.

Done when: the sentence names one reader and every claim has a named source.

## 2. Pick a rung

Four formats, one per rung. Each rung costs more to build and saves the reader
more effort.

| Rung | Format | The reader can | Cost |
|---|---|---|---|
| 1 | prose | read one idea once | minutes |
| 2 | diagram | see shape, flow, and nesting | minutes |
| 3 | HTML page | explore, compare, and open details | an afternoon |
| 4 | video | watch state change over time | a day or more |

Start from what the reader must do:

| The reader must | Rung |
|---|---|
| hold one idea after one read | 1 |
| see flow, structure, or nesting | 2 |
| move a number, compare options, or open a side path | 3 |
| watch cause and effect, or before and after | 4 |

The default stop is rung 3. Climb a rung when chat text failed. Drop a rung
when a diagram would carry one label. Rung 4 wins only when motion carries the
idea. If a still frame explains it, the page is cheaper. When the user names a
format, take it and skip the table.

To move an existing artifact to another rung, reuse its outline, diagrams, and
prose. Do not re-research.

Done when: the chosen rung matches one row of the second table.

## 3. Write controlled prose

Every rung carries prose. ASD-STE100 is the controlled-language spec written
for aerospace maintenance. Hold the writing to about 80 percent of it:

- One idea per sentence. Twelve words is normal, twenty is the ceiling.
- Active voice, present tense. "The window draws the frame", not "the frame
  is drawn".
- One term per concept, used every time. Name a thing once, then never call it
  by another word.
- Define a term in the sentence where it first appears.
- Fact first, then reason, then exception.
- Numbers over adjectives. "3.1 ms" beats "fast".
- No em dashes.

Before: "The rendering subsystem leverages an intelligent caching stratum to
eliminate redundant parse operations."

After: "The renderer parses the HTML once. Later drags reuse the parsed tree."

Done when: no sentence needs a second read.

## 4. Build the artifact

Rungs 1 and 2 need no file. Shape the prose, then draw:

- One idea per diagram. Split the idea instead of adding to the drawing.
- Label every part with the name the code or domain uses, so the diagram
  doubles as a map.
- One caption that states the idea.
- No stock art, no shadows. A picture that carries no meaning goes.

Rung 3 builds one HTML file. Deep recipe: `skills/html-report/SKILL.md` in
this repo.

- One `.html` file with inline CSS and JS. No CDN and no network request at
  any point, so it opens from `file://`.
- Light and dark themes, print styles, and a 60 to 70 character measure.
- Interactive where it teaches: `<details>` for side paths, tabs for
  alternatives, one control for a parameter. The first render reads with
  JavaScript off.
- A tl;dr of at most five lines at the top.

Rung 4 builds a video. Deep recipe:
`skills/html-report/reference/explainer-video.md` in this repo.

- 5 to 12 beats. One idea and one visual change per beat.
- Narration runs about 140 words per minute. Keep one sentence per audio file.
- Read the key from the environment and stop with a clear message when it is
  missing. Use ElevenLabs when `ELEVENLABS_API_KEY` or `XI_API_KEY` is set.
  Otherwise check `command -v` for `piper`, `coqui-tts`, `edge-tts`, or
  `espeak-ng`.
- Render with Manim for math, Remotion for React scenes, or HTML and ffmpeg
  for diagrams. Set each scene's duration from its audio.
- One `build.sh` renders from scratch. Running it twice produces the same
  video.

Done when: every claim from step 1 appears in the artifact, and no section is
a source inventory.

## 5. Check it, then let it go

- Read the artifact as the one reader. The one-line question must be
  answerable from it alone.
- Page: open it from `file://` and confirm zero network requests and a clean
  console.
- Video: watch it at 1x, start to finish. The first 15 seconds state the
  question, and the last frame names the takeaway.
- Run the unslop pass over every sentence of copy.
- Reply with the artifact's path, the one-line question, and the source count.
  Do not paste the artifact into chat.

Done when: the checks ran against the real artifact, not a summary of it.

## Why this matters

As models do more of the work, more human work becomes oversight and
understanding. Custom, discardable artifacts help there. A page or a video
that would never have paid for itself before now costs an afternoon. Ask for
the format that makes understanding cheapest.
