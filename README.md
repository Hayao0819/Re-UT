## Mozc Re-UT Dictionary

[本家Mozc-UT辞書様](http://linuxplayers.g1.xrea.com/mozc-ut.html)の代替を目標としたオープンソースプロジェクトです。

このプロジェクトはまだ非常に実験的で、うまくいくかわかりません。

### Todo

#### 二次配布OK

- [ ] alt-cannadic
- [ ] chimei
- [x] neologd
- [x] skkdic
- [ ] jinmei-ut
- [ ] sudachi dict

#### 二次配布NG

- [ ] jawiki-titles
- [ ] jawiki-articles
- [x] edict2

上記3つはGPLとの互換性がないため、他の辞書と同梱して配布することができません

- [ ] niconico IME  ~~(実装しても法的な問題がないかどうか確認中です)~~
- [ ] dic-nico-intersection-pixiv ~~(niconico IMEと同じく法的な問題を確認中)~~


上記2つは変換プログラムの実装には問題はなさそうです。ただし辞書の二次配布は明確に禁止されていました。

### License

このリポジトリに含まれているソースコードはWTFPLです。

生成された辞書は各配布元のライセンスに従ってください。

#### なぜWTFPLなのか

以下の理由からWTFPLを採用しました。

- 私が知っている限り最も自由なライセンス
- 一部のプロジェクトで採用実績があり、知名度も高い
- GitHubのライセンス表記でサポートされている
- GNU.orgによってGPL互換が[明記されている](https://www.gnu.org/licenses/license-list.html#WTFPL)
