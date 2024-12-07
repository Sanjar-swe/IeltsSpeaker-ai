# LinguaSpeak: AI Language Learning Assistant

## 🚀 О проекте

LinguaSpeak - это современное веб-приложение для изучения русского языка, использующее искусственный интеллект для улучшения процесса обучения.

## 🛠 Технологии

- Frontend: Next.js 14 + TypeScript + Tailwind CSS
- Backend: Golang
- AI: ElevenLabs API
- База данных: SQLite

## 📦 Установка и запуск

### Предварительные требования

- Node.js 18+
- Go 1.21+
- Docker и Docker Compose
- API ключ ElevenLabs

### Запуск проекта

1. Клонируйте репозиторий:
\`\`\`bash
git clone [url-репозитория]
cd linguaspeak
\`\`\`

2. Настройка Frontend:
\`\`\`bash
cd frontend
npm install
npm run dev
\`\`\`

3. Настройка Backend:
\`\`\`bash
cd backend
go mod download
go run cmd/server/main.go
\`\`\`

4. Или используйте Docker Compose:
\`\`\`bash
docker-compose up --build
\`\`\`

## 🔑 Переменные окружения

Создайте файл .env в корневой директории:

\`\`\`
ELEVENLABS_API_KEY=ваш_ключ
\`\`\`

## 📝 Функциональность

- Интерактивный чат с AI
- Определение и исправление грамматических ошибок
- Голосовое взаимодействие
- Культурные инсайты

## 👥 Разработка

Проект следует принципам чистой архитектуры и современным практикам разработки.
