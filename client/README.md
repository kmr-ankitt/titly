# Titly Client

Frontend for the Titly URL shortener, built with Svelte 5, SvelteKit, and Tailwind CSS.

## Features

- **Shortener Form**: URL validation with Zod, protocol auto-fixing, and error feedback for failed server requests.
- **Dynamic Routing (`/[url]`)**: Resolves `http://localhost:5173/:shortCode` against backend redirect endpoints with fallback 404 screens.
- **API Status**: Live backend server health badge in the navbar.
- **Link History**: Saves recent shortened URLs in `localStorage`.
- **Copy & QR Code**: Copy links to clipboard with toast notifications and generate QR codes.

## Project Structure

```text
src/
├── components/
│   ├── Form.svelte
│   ├── Output.svelte
│   ├── History.svelte
│   ├── Navbar.svelte
│   └── Toast.svelte
├── lib/
│   ├── api.ts
│   ├── history.ts
│   └── toastStore.ts
└── routes/
    ├── +layout.svelte
    ├── +page.svelte
    ├── +error.svelte
    └── [url]/
        ├── +page.ts
        └── +page.svelte
```

## Development & Build

```bash
# Install dependencies
pnpm install

# Development server (http://localhost:5173)
pnpm dev

# Type check
pnpm check

# Production build
pnpm build
```

## Environment Variables

| Variable | Default | Description |
| --- | --- | --- |
| `VITE_API_URL` | `http://localhost:4000` | Titly backend API base URL |
