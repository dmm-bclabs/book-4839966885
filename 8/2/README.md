# 8章フロントエンドのサンプルコード

## 概要

8章のフロントエンドのサンプルコードのDockerコンテナの構築、起動を行います。

### 実行手順

```
$ docker build . -t frontend
$ docker run -p 3000:3000 -it frontend
```
