もちろん、React＋Vite のフロントエンドに **Go（Golang）で作成したバックエンド API** を追加する拡張仕様書をまとめます。  
実務でもそのまま使えるレベルで、API 設計・通信方式・ディレクトリ構成まで整理しました。

---

# React＋Vite × Go API 拡張仕様書

## 1. 概要

### 1.1 目的  
既存の「カウンター＋ToDoリスト」アプリに、Go 言語で構築したバックエンド API を追加し、  
**フロントエンドとバックエンド間でデータ通信ができる構成**に拡張する。

### 1.2 ゴール  
- ToDo データを Go API で管理  
- React から API を呼び出して ToDo を取得・追加・削除  
- 将来的に DB（SQLite / PostgreSQL）へ拡張可能な構造にする

---

# 2. バックエンド（Go）仕様

## 2.1 使用技術
| 項目 | 内容 |
|------|------|
| 言語 | Go 1.22+ |
| Webフレームワーク | net/http（標準）または chi / gin（任意） |
| データ管理 | メモリ管理（初期）→ DB へ拡張可能 |
| API形式 | REST API |
| CORS | React（http://localhost:5173）からのアクセスを許可 |

---

## 2.2 API エンドポイント仕様

### 2.2.1 ToDo 一覧取得  
**GET /api/todos**

**レスポンス例**
```json
[
  { "id": 1, "title": "買い物に行く" },
  { "id": 2, "title": "メール返信" }
]
```

---

### 2.2.2 ToDo 追加  
**POST /api/todos**

**リクエスト**
```json
{ "title": "新しいタスク" }
```

**レスポンス**
```json
{ "id": 3, "title": "新しいタスク" }
```

---

### 2.2.3 ToDo 削除  
**DELETE /api/todos/{id}**

**レスポンス**
```json
{ "message": "deleted" }
```

---

## 2.3 Go 側のディレクトリ構成（推奨）

```
backend/
├─ main.go
├─ handlers/
│   └─ todo.go
├─ models/
│   └─ todo.go
└─ store/
    └─ memory.go
```

---

## 2.4 Go 側の処理フロー

1. サーバー起動（ポート例：8080）
2. メモリ上に ToDo の配列を保持
3. ルーティング設定
4. CORS 設定（React からのアクセス許可）
5. JSON の入出力を行うハンドラを実装

---

# 3. フロントエンド（React）拡張仕様

## 3.1 追加要件
| No | 内容 |
|----|------|
| F-01 | 初回ロード時に API から ToDo を取得 |
| F-02 | ToDo 追加時に API へ POST |
| F-03 | ToDo 削除時に API へ DELETE |
| F-04 | API エラー時の簡易エラーメッセージ表示 |

---

## 3.2 API 通信仕様（フロント側）

### 3.2.1 ToDo 取得
```js
fetch("http://localhost:8080/api/todos")
  .then(res => res.json())
  .then(data => setTodos(data));
```

### 3.2.2 ToDo 追加
```js
fetch("http://localhost:8080/api/todos", {
  method: "POST",
  headers: { "Content-Type": "application/json" },
  body: JSON.stringify({ title: input })
});
```

### 3.2.3 ToDo 削除
```js
fetch(`http://localhost:8080/api/todos/${id}`, {
  method: "DELETE"
});
```

---

# 4. CORS 設定

React（Vite）はデフォルトで  
**http://localhost:5173**  
で動作するため、Go 側で CORS を許可する。

例（標準 net/http の場合）：

- `Access-Control-Allow-Origin: *` または `http://localhost:5173`
- `Access-Control-Allow-Methods: GET, POST, DELETE`
- `Access-Control-Allow-Headers: Content-Type`

---

# 5. システム構成図

```
┌──────────────┐        HTTP (REST API)        ┌──────────────┐
│ React (Vite) │  <-------------------------->  │ Go API Server │
│  Frontend    │                                │  Backend      │
└──────────────┘                                └──────────────┘
```

---

# 6. 今後の拡張案

- DB（SQLite / PostgreSQL）に保存  
- JWT 認証の追加  
- Docker 化してコンテナで統合  
- WebSocket によるリアルタイム更新  
- Clean Architecture への発展  

---

必要であれば、  
**Go の実際のコード（main.go / handler / model）**  
や  
**React 側の API 通信コードを組み込んだ App.jsx**  
も作成できます。

どのレベルまで実装したいか教えてくれたら、そこまで一気に仕上げますよ。
