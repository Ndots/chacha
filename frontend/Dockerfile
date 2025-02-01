# Use the official Node.js image.
FROM node:18

WORKDIR /app

COPY package.json tsconfig.json next.config.js ./
COPY pages ./pages
COPY components ./components
COPY styles ./styles
COPY utils ./utils
COPY services ./services
COPY tailwind.config.js ./

RUN npm install

RUN npm run build

EXPOSE 3000

CMD ["npm", "start"]
