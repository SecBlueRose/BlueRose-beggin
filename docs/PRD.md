# Product Requirement Document (PRD)

## Project Name: Aegis
**Owner:** Khan Konstantin
**Status:** Approved
**Version:** 1.0 (MVP)

---

## 1. Problem Statement
Студентам и начинающим DevSecOps-инженерам не хватает единой, легкой и наглядной платформы для практики процессов ИБ: безопасной передачи секретов, мониторинга инфраструктуры и анализа уязвимостей в CI/CD без развертывания тяжеловесных enterprise-систем.

## 2. Product Vision & Goals
Создать модульную экосистему **Aegis**, объединяющую CLI-инструменты, microservices backend, автоматизированные сканеры безопасности и плотный веб-дашборд в стиле **JetBrains New UI**.

### Success Criteria for MVP (v0.1)
1. Работающий CLI-клиент на Go для взаимодействия с системой из терминала.
2. Безопасное хранение локальных секретов с TTL (Aegis Vault).
3. Мониторинг системных метрик и логов контейнеров (Aegis Monitor).
4. Автоматизированная проверка кода на утечки секретов в Git (Aegis Shield).

---

## 3. User Persona
* **Primary User:** DevSecOps / Cybersecurity Engineer.
* **Key Needs:** Скорость взаимодействия через CLI, минимализм, наглядность данных, поддержка стандартов DevSecOps.

---

## 4. Functional Requirements (Scope)

| ID       | Module        | Feature                                                | Priority |
|:---------|:--------------|:-------------------------------------------------------|:---------|
| **FR-1** | Core API      | REST API сервер для обработки команд CLI и Web UI      | **P0**   |
| **FR-2** | CLI Client    | Консольная утилита `aegis-cli`                         | **P0**   |
| **FR-3** | Aegis Vault   | Эндпоинты зашифрованного сохранения и вычитки секретов | **P0**   |
| **FR-4** | Aegis Monitor | Сбор статуса Docker-контейнеров и логов                | **P1**   |
| **FR-5** | Aegis Shield  | Сканирование коммитов (Pre-commit hook / CI/CD)        | **P1**   |
| **FR-6** | Web UI        | Демо-панель в стиле JetBrains IDE                      | **P2**   |

---

## 5. Non-Functional Requirements
* **Architecture:** Монорепозиторий с микросервисной архитектурой.
* **Deployment:** Развертывание всей среды одной командой через Docker Compose.