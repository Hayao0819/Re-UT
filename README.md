## Mozc Re-UT Dictionary

[本家Mozc-UT辞書様](http://linuxplayers.g1.xrea.com/mozc-ut.html)の代替を目標としたオープンソースプロジェクトです。

このプロジェクトはまだ非常に実験的で、うまくいくかわかりません。

### 現状

基本的な機能を実装している段階であり、現在neologdのみ実装されています。

`src/<各辞書>/run.sh`で`build/<各辞書>/dic.txt`を生成し、それを最後に結合する仕組みに使用と考えています。

`src/make.sh`で各`run.sh`を実行するつもりです。

今はneologd辞書の膨大なCSVの変換の高速化を行おうとしています。

### License

このリポジトリに含まれているソースコードはWTFPLです。

生成された辞書は各配布元のライセンスに従ってください。

#### なぜWTFPLなのか

以下の理由からWTFPLを採用しました。

- 私が知っている限り最も自由なライセンス
- 一部のプロジェクトで採用実績があり、知名度も高い
- GitHubのライセンス表記でサポートされている
- GNU.orgによってGPL互換が[明記されている](https://www.gnu.org/licenses/license-list.html#WTFPL)
