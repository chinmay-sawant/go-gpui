# go-gpui website

A React and Vite site with the README demos and the repository's Markdown documentation. It uses plain CSS and native video controls.

## Run locally

Use Node.js 22.12 or newer.

```sh
cd frontend
npm ci
npm run dev
```

Vite prints the local address under `/go-gpui/`. `npm test` checks the GitHub cache. `npm run build` writes the static site to the repository's root `docs/` folder; `npm run preview` serves that build. Each build replaces the generated output in `docs/`.

## GitHub Pages

Run `npm run build` from `frontend/`. The build includes the videos, images, and `.nojekyll` file, and uses `/go-gpui/` for asset paths. Keep the generated `docs/` folder in the repository when publishing. In the repository's Settings > Pages, choose **Deploy from a branch**, select the publishing branch, and select **/docs**.

For a custom domain hosted at the root, build with `npm run build -- --base=/`.

## Content

The site imports `../documentation/*.md`, the example index, desktop cat guide, Android guide, and `../showcase.md` at build time. Edit those files to update the website. Relative documentation links stay in the site; source-code and plan links open the repository on GitHub.

The two videos in `public/demos/` are the original MP4s from the X posts linked in the README. Their poster images are the existing `assets/preview.webp` and `assets/desktop-cat.webp`. The site serves the videos locally and links to the original posts. Videos load metadata only and do not autoplay.

## GitHub stars

The top-right link reads `stargazers_count` from GitHub's public repository API. Successful counts are cached in localStorage for six hours. Failed requests back off for at least one hour and retain the last count. The loader honors longer `Retry-After` and rate-limit reset times, limits requests to ten seconds, and shares an in-flight request within the page. Web Locks coordinate tabs where supported. If browser storage is blocked, the page still uses its memory cache.

This cache is per browser, not shared across visitors. There is no polling or client-side token. A shared cache at the hosting layer can be added when deployment is chosen.

API reference: [GitHub repository endpoint](https://docs.github.com/en/rest/repos/repos#get-a-repository). Build setup: [Vite documentation](https://vite.dev/guide/).
