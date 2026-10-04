# Explainer video

Branch of `html-report`. The user asks for a video or runs
`/html-report --video <topic>`.

Video wins when motion carries the idea: state over time, cause and effect,
before and after. When the topic is static, a page wins. Say that to the
user before spending render time.

If the user names a style ("3b1b style"), keep it. That style means a dark
background, Manim, a narrator who says where the viewer stands, and
equations built term by term.

## 1. Beat sheet

Five to twelve beats, one idea per beat. Each beat gets one narration
sentence and one sentence on what changes on screen. No beat without a
visual change.

Done when: every beat earns its screen time, and the last one states the
takeaway.

## 2. Script

Use the controlled-language list from the page. Narration runs about 140
words per minute, so a 4 minute video is about 560 words. Keep one sentence
per audio file.

Done when: the word count fits the target duration and every sentence passes
the read-once test.

## 3. Narration

Never hardcode a key. Read it from the environment and stop with a clear
message when it is missing.

- ElevenLabs, when `ELEVENLABS_API_KEY` or `XI_API_KEY` is set:

```bash
curl -s -X POST "https://api.elevenlabs.io/v1/text-to-speech/$VOICE_ID" \
  -H "xi-api-key: $ELEVENLABS_API_KEY" -H "Content-Type: application/json" \
  -d '{"text":"LINE","model_id":"eleven_multilingual_v2"}' -o line-01.mp3
```

- Local options, when no key is set. Check with `command -v`: `piper` (local,
  fast), `coqui-tts` (local, heavier), `edge-tts` (free, needs network),
  `espeak-ng` (robotic fallback).
- Cache one audio file per line. A re-render skips lines whose text did not
  change.

Done when: every line has an audio file and `ffprobe` reports a non-zero
duration.

## 4. Visuals

Pick the smallest tool that covers the beats:

- Manim for math and 3b1b style.
- Remotion for React-driven scenes, charts, and typography.
- HTML plus Playwright screenshots plus ffmpeg for diagrams and text.
- Real data beats fake animation. Only animate numbers a script computed.

Done when: every beat has a scene, and a still of each scene reads as
clearly as the page's diagrams.

## 5. Assemble

- Set each scene's duration from its audio file:
  `ffprobe -show_entries format=duration`.
- Concat the audio tracks in beat order, then mux them against the rendered
  scenes.
- Add captions from the script. Keep them on by default; many people watch
  muted.
- One `build.sh` renders from scratch. Running it twice produces the same
  video.

Done when: `build.sh` runs clean twice in a row.

## 6. Check it

- Watch the whole video at 1x, start to finish. No skipping.
- Confirm the first 15 seconds state the question.
- Confirm each scene starts and ends inside its narration, with no dead air
  and no clip.
- Confirm the last frame names the takeaway.

Done when: the full watch happened and all three checks hold.

## Deliverables

`plans/reports/video/<slug>/`: `beats.tsv`, `script.md`, `audio/`,
`scenes/`, `out.mp4`, `build.sh`.

The video is disposable too. The script outlives it.
