FROM node:22-alpine

WORKDIR /app

RUN npm install -g pnpm

COPY . .

RUN pnpm install --frozen-lockfile

RUN pnpm build

EXPOSE 5173

CMD ["pnpm", "preview", "--host", "0.0.0.0"]
