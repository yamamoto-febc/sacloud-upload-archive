# sacloud-upload-archive

## Overview

さくらのクラウドへアーカイブのアップロードを行うためのコマンドです。

## Description

さくらのクラウドには[アーカイブ](http://cloud.sakura.ad.jp/specification/server-disk/index.html#server-disk-content03)という機能があります。
通常さくらのクラウドでサーバを作成する際はあらかじめ初期設定済みのOSイメージ(アーカイブ)からOSを選択するのですが、自分でカスタマイズしたアーカイブをアップロードして好きなOSをインストールすることも可能です。

ただ、アーカイブのアップロードは割と煩雑な作業です。詳細は[こちらのページ](https://help.sakura.ad.jp/app/answers/detail/a_id/2522/c/342)に記載されているのですが、

    1) コントロールパネル上でアーカイブ用領域を作成
    2) ファイル領域へのFTP接続情報が表示されるため控えておく
    3) FTPS+PASV対応のFTPクライアントからアップロード

という手順が必要になります。

`sacloud-upload-archive`を使えばコマンド一発でアップロードが行えるようになります。

## Install

こちらの[リリースページ](https://github.com/yamamoto-febc/sacloud-upload-archive/releases/latest)から最新バージョンのバイナリをダウンロードして展開、実行権を付与してください。

以下のプラットフォーム用のバイナリを用意しています。
  * darwin(i386/amd64)
  * linux(i386/amd64)
  * windows(i386/amd64)

## Usage

オプション、イメージのパス、イメージ名を指定するとアップロード実施します。
コマンドの書式は以下の通りです。

```bash

$ sacloud-upload-archive [オプション] [イメージ名]

```

#### オプション

    -token  : さくらのクラウドのAPIキー(アクセストークン)
    -secret : さくらのクラウドのAPIキー(シークレット)
    -zone   : 作成するゾーン (is1a/is1b/tk1a) デフォルト:is1a
    -file   : アップロードするアーカイブのファイルパス

`zone`は以下の値をとります
>
    is1a : 石狩第１ゾーン
    is1b : 石狩第２ゾーン
    tk1a : 東京第１ゾーン

#### APIキーの設定方法

上記のオプションで指定する、または環境変数も利用できます。

* オプションで指定する場合

```オプションで指定する場合
$ sacloud-upload-archive -token=[アクセストークン] -secret=[シークレット] [イメージ名]
```

* 環境変数で指定する場合

```環境変数で指定する場合
$ export SAKURACLOUD_ACCESS_TOKEN=[アクセストークン]
$ export SAKURACLOUD_ACCESS_TOKEN_SECRET=[シークレット]
$ sacloud-upload-archive [イメージ名]
```

#### コマンドでのアーカイブ指定方法

アーカイブは`-file`オプション、またはパイプ/リダイレクトで指定します。
`curl`などでアーカイブを取得、パイプで渡してさくらのクラウドへアップロードという使い方ができます。

```アーカイブ指定方法の例

# -fileオプションで指定(config2.isoというファイルをアップロードする例)
$ sacloud-upload-archive [オプション] -file ./config2.iso [イメージ名] 

# curlからパイプで受け取る
$ curl -L http://[ISOイメージのURL] | sacloud-upload-archive [オプション] [イメージ名]

# リダイレクトで受け取る(config2.isoというファイルをアップロードする例)
$ sacloud-upload-archive [オプション] [イメージ名] < config2.iso

```


## License

This project is published under [Apache 2.0 License](LICENSE).

## Author

* Kazumichi Yamamoto ([@yamamoto-febc](https://github.com/yamamoto-febc))
