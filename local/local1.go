package local

// 小文字で始まるとパッケージ内のみ：同一パッケージからしか呼べません.
var localStr = "Localだよ" // <- Package local 内でのみ参照可能
// 大文字で始まるとパッケージ外参照可能.
var GlobalStr = "Package localのGlobalStrという変数だよ" // <- Package 外に公開
