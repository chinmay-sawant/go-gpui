package remote

// darkTheme is the default. lightTheme is the other side of the toggle.
// Both only set the variables the template reads, so the layout stays put.
const darkTheme = `
:root {
  --bg: #101010;
  --ink: #f2f2f2;
  --muted: #9a9a9a;
  --key: #242426;
  --line: #343438;
  --power: #c43838;
  --pad: #1a1a1f;
}
`

const lightTheme = `
:root {
  --bg: #f3f1ec;
  --ink: #1a1a1a;
  --muted: #6d675f;
  --key: #e7e2d8;
  --line: #d5d0c6;
  --power: #a32020;
  --pad: #ffffff;
}
`
