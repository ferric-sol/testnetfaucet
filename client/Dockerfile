FROM node:20 AS builder

WORKDIR /app

# Copy only files needed for installing dependencies
COPY package.json yarn.lock ./

# Install dependencies (frozen-lockfile ensures exact versions)
RUN yarn install --frozen-lockfile

# Copy the rest of the application source
COPY . .

# Build the Next.js app for production
RUN yarn build

# Final image or Runtime Stage: Create lightweight image for running the app
FROM node:20-alpine AS runner

WORKDIR /app

# Copy the compiled binaries from build
COPY --from=builder /app ./

# Optionally add .env.production if using ENV vars
# COPY --from=builder /app/.env.production ./.env.production

EXPOSE 3000

# Run command
CMD ["yarn", "start"]
