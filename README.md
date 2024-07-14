# URLShortner

Rewritten from TypeScript to Go.

## Why the $?! would you create Base68 URL encoding... with unsafe characters?

Because I can?

Really I needed a URL shortner for one of my personal projects I. I decided it would be fun to write one that was an unusal base for I had to roll my own base system and handle the incrementing of it. So I picked 68 (mostly) safe characters.

You can read all about [URI reserved characters](https://www.rfc-editor.org/rfc/rfc3986#section-2.2). The reserved characters I decided to use are sub-delims. I am already using all of the unreserved characters as spcified in [URI Unreserved Characters](https://www.rfc-editor.org/rfc/rfc3986#section-2.3)

So really this insanity is just me having fun.
