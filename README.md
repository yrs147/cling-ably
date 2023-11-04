# Cling 
![cling-removebg-preview (1)](https://github.com/yrs147/cling-ably/assets/98258627/be18c246-db29-4a01-b494-661df2863fdd)

## Screenshot

<img width="1720" alt="Screenshot 2023-11-02 at 1 40 13 AM" src="https://github.com/yrs147/cling-ably/assets/98258627/706cdb18-b60c-40f8-8449-689edd861e90">



## Overview
Cling is a real-time CLI chat application designed to bridge language barriers by connecting people from different parts of the world and enabling them to collaborate without worrying about the language barrier. It leverages the power of the OpenNMT ArgoTranslate machine-learning model for real-time message translation and the `Ably Realtime Framework` for a seamless communication experience.

### Supported Languages
Cling supports the following languages :
- **English (en)**
- **Hindi (hi)**
- **French (fr)**
- **German (de)**
- **Italian (it)**
- **Polish (pl)**
- **Spanish (es)**
-  **Russian (ru)**
-  **Spanish (es)**
-  **Japanese (ja)**
-  **Chinese (zh)**

## Prerequisites
Before you get started with Cling, ensure you have the following prerequisites installed:
- [ABLY_KEY](https://ably.com/)
- [Go](https://golang.org/)
- [Docker](https://www.docker.com/get-started)

## Setup
Follow these steps to set up the Cling project

1. Run the LibreTranslate Docker container to load supported languages
```
docker run -it -p 5000:5000 libretranslate/libretranslate --load-only en,hi,fr,de,it,pl,es,ru,ja,zh
```
2. Clone the Cling repository

```
git clone https://github.com/yrs147/cling-ably.git
cd cling-ably
```
3. Get your `ABLY_KEY` and add it to the `.env` file

4. Add the Dependencies
```
go mod tidy

```

5. Build the binary
```
go build -o cling
```

6. Run Cling

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

