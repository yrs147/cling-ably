# Cling 
![cling-removebg-preview (1)](https://github.com/yrs147/cling-ably/assets/98258627/be18c246-db29-4a01-b494-661df2863fdd)



## Overview
Cling is a real-time CLI chat application designed to bridge language barriers by connecting people from different parts of the world. It leverages the power of the OpenNMT ArgoTranslate machine learning model for real-time message translation. Cling is built on the `Ably Realtime Framework`.

### Supported Languages
Cling supports the following languages with their language codes:
- **English (en)**
- **Hindi (hi)**
- **French (fr)**
- **German (de)**
- **Italian (it)**
- **Polish (pl)**
- **Spanish (es)**
-  **Russian (ru)**
-  **Spanish (es)**
-  **Japanese (js)**
-  **Chinese (zh)**

## Prerequisites
Before you get started with Cling, ensure you have the following prerequisites installed:
- [ABLY_KEY](https://ably.com/)
- [Go](https://golang.org/)
- [Docker](https://www.docker.com/get-started)

## Setup
Follow these steps to set up the Cling project:

1. Run the LibreTranslate Docker container to load supported languages:
   ```bash
   docker run -it -p 5000:5000 libretranslate/libretranslate --load-only en,hi,fr,de,it,pl,es,ru
   ```
2. Clone the Cling repository:

```
git clone <repo_link>
cd cling
```
3. Add the Dependencies
```
go mod tidy

```

4. Build the binary
```
go build -o cling
```

5. Run Cling

```
./cling -u <username> -r <roomname> -l <preferredlanguage>

```

## Command and Flags

Cling offers the following command and flags:

```
Usage: cling [flags]

A CLI chat app using Ably

Flags:
  -h, --help          Help for cling
  -u, --username      Your username for the chat (default "defaultUsername")
  -r, --room          Chat room code (default "defaultRoomCode")
  -l, --language      Chat room language code (default "en")

```

- `username (-u)`: Set your username for the chat (default is "defaultUsername").
- `room (-r)`: Set the chat room code (default is "defaultRoomCode").
- `language (-l)`: Set the chat room language code (default is "en").

