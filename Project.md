LinguaSpeak: AI Language Learning Assistant
🚀 Проект: Веб-приложение для изучения русского языка с AI
🔧 Технологический стек

Frontend: Next.js 14 + TypeScript
Styling: Tailwind CSS
Backend: Golang
AI Integration: ElevenLabs API
Database: SQLite

📂 Структура проекта
Copylinguaspeak/
│
├── frontend/ # Next.js Frontend
│ ├── components/
│ │ ├── ChatInterface/
│ │ ├── AuthForm/
│ │ ├── VoiceRecorder/
│ │ └── ErrorHighlighter/
│ │
│ ├── pages/
│ │ ├── index.tsx # Главная страница
│ │ ├── chat.tsx # Страница чата
│ │ └── profile.tsx # Страница профиля
│ │
│ ├── styles/
│ │ └── globals.css # Глобальные стили Tailwind
│ │
│ ├── hooks/ # Custom React hooks
│ │ ├── useChat.ts
│ │ └── useVoiceRecorder.ts
│ │
│ └── utils/ # Утилиты
│ ├── errorHighlight.ts
│ └── apiClient.ts
│
├── backend/ # Golang Backend
│ ├── cmd/
│ │ └── server/
│ │ └── main.go # Точка входа сервера
│ │
│ ├── internal/
│ │ ├── handlers/ # HTTP обработчики
│ │ ├── models/ # Модели данных
│ │ └── services/ # Бизнес-логика
│ │
│ └── pkg/ # Общие утилиты
│ ├── ai/ # AI интеграции
│ └── database/ # Работа с БД
│
└── docker-compose.yml # Конфигурация развертывания
🎯 Ключевые компоненты для генерации

1. Frontend Components
   ChatInterface

Создать интерактивный чат-компонент
Поддержка текстовых и голосовых сообщений
Визуализация грамматических ошибок

typescriptCopy// Примерная структура компонента
interface ChatMessage {
id: string;
text: string;
sender: 'user' | 'ai';
errors?: {
word: string;
correction: string;
}[];
}

function ChatInterface() {
// Логика чата с AI
}
ErrorHighlighter

Функция определения и подсветки грамматических ошибок
Генерация корректировок

typescriptCopyfunction highlightErrors(text: string): {
highlightedText: React.ReactNode;
errors: Array<{word: string, correction: string}>;
} {
// Логика определения ошибок
} 2. Backend Services
AI Error Detection Service

Микросервис для анализа текста
Интеграция с языковыми AI моделями

goCopytype ErrorDetectionService struct {
// Настройки и методы для обработки текста
DetectGrammarErrors(text string) []GrammarError
}
🔍 Основные задачи генерации кода

Создание компонентов frontend
Разработка backend-логики
Настройка интеграции с ElevenLabs API
Реализация механизма обработки ошибок
Создание системы аутентификации

🛠 Инструкции для Windsurf AI IDE
Генерация Frontend

Использовать TypeScript
Применять Tailwind для стилизации
Фокус на компонентной архитектуре
Реализовать адаптивный дизайн

Генерация Backend

Использовать чистый Golang
Применять принципы Clean Architecture
Создавать тестируемый код
Обеспечивать безопасность и производительность

📋 Промпт для генерации кода
CopyСоздай компонент ChatInterface для Next.js:

- Поддержка текста и голоса
- Визуализация грамматических ошибок
- Интеграция с ElevenLabs API
- Использовать TypeScript и Tailwind
  🚦 Этапы разработки

Базовая структура проекта
UI/UX компоненты
AI-интеграция
Тестирование
Деплой

📦 Зависимости

nextjs
tailwindcss
typescript
elevenlabs-api
zod (валидация)
next-auth (аутентификация)

🔒 Безопасность

Валидация входных данных
Безопасная аутентификация
Защита от XSS
Лимитирование запросов

🌐 Деплой

Vercel для Frontend
Railway или Heroku для Backend
Docker для контейнеризации
