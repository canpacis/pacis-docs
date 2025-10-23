# Assets

Assets are the non-Go files used throughout your application such as images, scripts, and stylesheets.
Some assets can be embedded directly into the final binary, while others need to be transformed for performance and caching purposes.

## Web Assets

Your JavaScript, TypeScript, and CSS files are bundled using **Vite**.
These files are part of your codebase and live under `src/web/`.

The Pacis template includes a few example files:

```
src/web/
├── main.ts
├── stream.ts
└── style.css
```

You reference these files from your Go server using a special function.
During development, Pacis proxies Vite’s dev server to enable **hot module replacement (HMR)**.
In production, it automatically resolves the correct paths to the bundled assets.

{noaccessory=true}
```go
Script(Type("module"), Src(server.Asset("/src/web/main.ts")))
```

> Note: The `server.Asset` function simply returns a string.

When you build for production, these files are transformed and renamed based on a **content hash**.
This allows browsers to cache unchanged files aggressively, while automatically invalidating caches when the file content changes.

Example output:

```
https://yourdomain.com/style.CX7LkpoD.css
https://yourdomain.com/stream.BGrupRFG.js
```

## Static Assets

Static assets, such as images, icons, or plain text files, can be placed in the `public/` directory and referenced directly.

Example structure:

```
└── root/
    └── public/
        └── background.png
```

You can use the image in your templates like so:

{noaccessory=true}
```go
Img(Src("/background.png"))
```

When compiled, files inside `public/` will be available at the root of your application with their original names.
This makes it a good place for files like `robots.txt`, `sitemap.xml`, or `favicon.ico`.

Example URLs:

```
https://yourdomain.com/background.png
https://yourdomain.com/robots.txt
https://yourdomain.com/sitemap.xml
https://yourdomain.com/favicon.ico
```
