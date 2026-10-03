package theme

// lightTheme and darkTheme are extra stylesheets. Each sets the custom
// properties the template reads with var(). Both keep the same geometry --
// same padding, same box sizes -- so switching moves nothing; only the paint
// changes. SetTheme applies one after the template's own styles, so a theme
// rule wins a tie.
const lightTheme = `
:root {
  --bg: #f4f1ea;
  --card: #ffffff;
  --ink: #1c1915;
  --accent: #1a56db;
  --radius: 10px;
}
`

const darkTheme = `
:root {
  --bg: #14161a;
  --card: #1e2228;
  --ink: #e8e6e1;
  --accent: #7aa2f7;
  --radius: 18px;
}
`
