|kinds of single-character expressions|单字符匹配	|examples|
|:-|:-|:-|
|any character, possibly including newline (s=true)|匹配全部字符,当s=true时,可以匹配换行符	|.|
|character class|匹配"字符类"中的一个字符|	[xyz]|
|negated character class|匹配"字符类"以外的一个字符	|[^xyz]|
|Perl character class (link)|匹配"Perl字符类"中的一个字符|	\d|
|negated Perl character class|匹配"Perl字符类"以外的一个字符|	\D|
|ASCII character class (link)|匹配"ACSII字符类"中的一个字符|	[[:alpha:]]|
|negated ASCII character class|匹配"ACSII字符类"以外的一个字符|	[[:^alpha:]]|
|Unicode character class (one-letter name)|匹配"Unicode字符类"中的一个字符(仅一个字母)|	\pN|
|Unicode character class	|匹配"Unicode字符类"中的一个字符|\p{Greek}|
|negated Unicode character class (one-letter name)|匹配"Unicode字符类"以外的一个字符(仅一个字母)|	\PN|
|negated Unicode character class	|匹配"Unicode字符类"以外的一个字符|\P{Greek}|

|expression|Composites|混合型|
|:-|:-|:-|
|xy|	x followed by y|匹配 xy（x 后面跟随 y）|
|x|y|	x or y (prefer x)|匹配 x 或 y (优先匹配 x)|

|expression|Repetitions|重复型|
|:-|:-|:-|
|x*|	zero or more x, prefer more|匹配零个或多个 x，优先匹配更多(贪婪)|
|x+|	one or more x, prefer more|匹配一个或多个 x，优先匹配更多(贪婪)|
|x?|	zero or one x, prefer one|匹配零个或一个 x，优先匹配一个(贪婪)|
|x{n,m}|	n or n+1 or ... or m x, prefer more|匹配 n 到 m 个 x，优先匹配更多(贪婪)|
|x{n,}|	n or more x, prefer more|匹配 n 个或多个 x，优先匹配更多(贪婪)|
|x{n}|	exactly n x|只匹配 n 个 x|
|x*?|	zero or more x, prefer fewer|匹配零个或多个 x，优先匹配更少(非贪婪)|
|x+?|	one or more x, prefer fewer|匹配一个或多个 x，优先匹配更少(非贪婪)|
|x??|	zero or one x, prefer zero|匹配零个或一个 x，优先匹配零个(非贪婪)|
|x{n,m}?|	n or n+1 or ... or m x, prefer fewer|匹配 n 到 m 个 x，优先匹配更少(非贪婪)|
|x{n,}?|	n or more x, prefer fewer|匹配 n 个或多个 x，优先匹配更少(非贪婪)|
|x{n}?|	exactly n x|只匹配 n 个 x|
|x{}|	(≡ x*) (NOT SUPPORTED) VIM|不支持VIM|
|x{-}|	(≡ x*?) (NOT SUPPORTED) VIM|不支持VIM|
|x{-n}|	(≡ x{n}?) (NOT SUPPORTED) VIM|不支持VIM|
|x=|	(≡ x?) (NOT SUPPORTED) VIM|不支持VIM|

Implementation restriction: The counting forms x{n,m}, x{n,}, and x{n} reject forms that create a minimum or maximum repetition count above 1000. Unlimited repetitions are not subject to this restriction.  
使用限制：计数形式x {n，m}，x {n，}和x {n}拒绝创建1000以上的最小或最大重复次数的表单。无限重复不受此限制。

|expression|Possessive repetitions|侵占重复型|
|:-|:-|:-|
|x*+|	zero or more x, possessive (NOT SUPPORTED)|匹配零个或多个 x,侵占(不支持)|
|x++|	one or more x, possessive (NOT SUPPORTED)|匹配一个或多个 x,侵占(不支持)|
|x?+|	zero or one x, possessive (NOT SUPPORTED)|匹配零个或一个 x,侵占(不支持)|
|x{n,m}+|	n or ... or m x, possessive (NOT SUPPORTED)|匹配 n 到 m 个 x,侵占(不支持)|
|x{n,}+|	n or more x, possessive (NOT SUPPORTED)|匹配 n 个或多个 x,侵占(不支持)|
|x{n}+|	exactly n x, possessive (NOT SUPPORTED)|只匹配 n 个 x,侵占(不支持)|

|expression|Grouping|分组|
|:-|:-|:-|
|(re)|	numbered capturing group (submatch)|编号捕获组(子匹配)|
|(?P<name>re)|	named & numbered capturing group (submatch)|命名和编号捕获组(子匹配)|
|(?<name>re)|	named & numbered capturing group (submatch) (NOT SUPPORTED)|命名和编号捕获组(子匹配)(不支持)|
|(?'name're)|	named & numbered capturing group (submatch) (NOT SUPPORTED)|命名和编号捕获组(子匹配)(不支持)|
|(?:re)|	non-capturing group|非捕获组|
|(?flags)|	set flags within current group; non-capturing|设置当前组内的标志;非捕获|
|(?flags:re)|	set flags during re; non-capturing|在re期间设置标志;非捕获|
|(?#text)|	comment (NOT SUPPORTED)|评论(不支持)|
|(?|x|y|z)|	branch numbering reset (NOT SUPPORTED)|分支编号复位(不支持)|
|(?>re)|	possessive match of re (NOT SUPPORTED)|所有格匹配(不支持)|
|re@>|	possessive match of re (NOT SUPPORTED) VIM|所有格匹配(不支持)VIM|
|%(re)|	non-capturing group (NOT SUPPORTED) VIM|非捕获组(不支持)VIM|

|expression|Flags|标记|
|:-|:-|:-|
|i|	case-insensitive (default false)|不区分大小写(默认是false)|
|m|	multi-line mode: ^ and $ match begin/end line in addition to begin/end text (default false)|多行模式：让 ^ 和 $ 匹配整个文本的开头和结尾，而非行首和行尾(默认为 false)|
|s|	let . match \n (default false)|让 . 匹配 \n (默认为 false)|
|U|	ungreedy: swap meaning of x* and x*?, x+ and x+?, etc (default false)|非贪婪模式：交换 x* 和 x*? 等的含义 (默认为 false)|
|Flag| syntax is xyz (set) or -xyz (clear) or xy-z (set xy, clear z).||

|expression|Empty strings|位置标记|
|:-|:-|:-|
|^|	at beginning of text or line (m=true)|如果标记 m=true 则匹配行首，否则匹配整个文本的开头（m 默认为 false）|
|$|	at end of text (like \z not \Z) or line (m=true)|如果标记 m=true 则匹配行尾，否则匹配整个文本的结尾（m 默认为 false）|
|\A|	at beginning of text|匹配整个文本的开头|
|\b|	at ASCII word boundary (\w on one side and \W, \A, or \z on the other)|匹配ASCII边界（\ w在一边，\ W，\ A或\ z在另一边）|
|\B|	not at ASCII word boundary|匹配非ASCII边界|
|\g|	at beginning of subtext being searched (NOT SUPPORTED) PCRE|匹配搜索子文本的开头（不支持）PCRE时|
|\G|	at end of last match (NOT SUPPORTED) PERL|在文本结尾的最后匹配|
|\Z|	at end of text, or before newline at end of text (NOT SUPPORTED)|匹配整个文本的结尾，或在文本结尾的换行符之前|
|\z|	at end of text|匹配整个文本的结尾，忽略 m 标记|
|(?=re)|	before text matching re (NOT SUPPORTED)|在文本匹配之前(不支持)|
|(?!re)|	before text not matching re (NOT SUPPORTED)|在文本不匹配之前(不支持)|
|(?<=re)|	after text matching re (NOT SUPPORTED)|在文本匹配之后(不支持)|
|(?<!re)|	after text not matching re (NOT SUPPORTED)|在文本不匹配之后(不支持)|
|re&|	before text matching re (NOT SUPPORTED) VIM|在文本匹配之前(不支持)VIM|
|re@=|	before text matching re (NOT SUPPORTED) VIM|在文本匹配之前(不支持)VIM|
|re@!|	before text not matching re (NOT SUPPORTED) VIM|在文本不匹配之前(不支持)VIM|
|re@<=|	after text matching re (NOT SUPPORTED) VIM|在文本匹配之后(不支持)VIM|
|re@<!|	after text not matching re (NOT SUPPORTED) VIM|在文本不匹配之后(不支持)VIM|
|\zs|	sets start of match (= \K) (NOT SUPPORTED) VIM|设置匹配开始(= \K)(不支持)VIM|
|\ze|	sets end of match (NOT SUPPORTED) VIM|设置匹配结束(不支持)VIM|
|\%^|	beginning of file (NOT SUPPORTED) VIM|文件开始(不支持)VIM|
|\%$|	end of file (NOT SUPPORTED) VIM|文件结束(不支持)VIM|
|\%V|	on screen (NOT SUPPORTED) VIM|在屏幕上(不支持)VIM|
|\%#|	cursor position (NOT SUPPORTED) VIM|光标位置(不支持)VIM|
|\%'m|	mark m position (NOT SUPPORTED) VIM|标记m的位置(不支持)VIM|
|\%23l|	in line 23 (NOT SUPPORTED) VIM|在23行(不支持)VIM|
|\%23c|	in column 23 (NOT SUPPORTED) VIM|在23栏(不支持)VIM|
|\%23v|	in virtual column 23 (NOT SUPPORTED) VIM|在虚拟栏23(不支持)VIM|

|expression|Escape sequences|转义序列|
|:-|:-|:-|
|\a|	bell (≡ \007)|匹配响铃符,(相当于\007)|
|\f|	form feed (≡ \014)|匹配换页符,(相当于\014)|
|\t|	horizontal tab (≡ \011)|匹配横向制表符,(相当于\011)|
|\n|	newline (≡ \012)|匹配换行符,(相当于\012)|
|\r|	carriage return (≡ \015)|匹配回车符,(相当于\015)|
|\v|	vertical tab character (≡ \013)|匹配纵向制表符,(相当于\013)|
|\*|	literal *, for any punctuation character *|文本*，用于任何标点符号*|
|\123|	octal character code (up to three digits)|匹配八进制字符代码(最多三位数字)|
|\x7F|	hex character code (exactly two digits)|匹配十六进制字符代码(正好两位数字)|
|\x{10FFFF}|	hex character code|匹配十六进制字符代码|
|\C|	match a single byte even in UTF-8 mode|UTF-8模式下也匹配单个字节|
|\Q...\E|	literal text ... even if ... has punctuation|匹配 \Q 和 \E 之间的文本，忽略文本中的正则语法|
|\1|	backreference (NOT SUPPORTED)|反向引用|
|\b|	backspace (NOT SUPPORTED) (use \010)|退格符|
|\cK|	control char ^K (NOT SUPPORTED) (use \001 etc)||
|\e|	escape (NOT SUPPORTED) (use \033)||
|\g1|	backreference (NOT SUPPORTED)||
|\g{1}|	backreference (NOT SUPPORTED)||
|\g{+1}|	backreference (NOT SUPPORTED)||
|\g{-1}|	backreference (NOT SUPPORTED)||
|\g{name}|	named backreference (NOT SUPPORTED)||
|\g<name>|	subroutine call (NOT SUPPORTED)||
|\g'name'|	subroutine call (NOT SUPPORTED)||
|\k<name>|	named backreference (NOT SUPPORTED)||
|\k'name'|	named backreference (NOT SUPPORTED)||
|\lX|	lowercase X (NOT SUPPORTED)||
|\ux|	uppercase x (NOT SUPPORTED)||
|\L...\E|	lowercase text ... (NOT SUPPORTED)||
|\K	reset| beginning of $0 (NOT SUPPORTED)||
|\N{name}|	named Unicode character (NOT SUPPORTED)||
|\R|	line break (NOT SUPPORTED)||
|\U...\E|	upper case text ... (NOT SUPPORTED)||
|\X|	extended Unicode sequence (NOT SUPPORTED)||
|\%d123|	decimal character 123 (NOT SUPPORTED) VIM||
|\%xFF|	hex character FF (NOT SUPPORTED) VIM||
|\%o123|	octal character 123 (NOT SUPPORTED) VIM||
|\%u1234|	Unicode character 0x1234 (NOT SUPPORTED) VIM||
|\%U12345678|	Unicode character 0x12345678 (NOT SUPPORTED) VIM||

|expression|Character class elements|字符类元素|
|:-|:-|:-|
|x|	single character|单个字符|
|A-Z|	character range (inclusive)|字符范围(含)|
|\d|	Perl character class|Perl匹配字符类|
|[:foo:]|	ASCII character class foo|ASCII字符类foo|
|\p{Foo}|	Unicode character class Foo|Unicode字符类foo|
|\pF|	Unicode character class F (one-letter name)|Unicode字符类F(单字母名称)|

|expression|Named character classes as character class elements|命名字符类作为字符类元素|
|:-|:-|:-|
|[\d]|	digits (≡ \d)|匹配数字 (相当于 \d)|
|[^\d]|	not digits (≡ \D)|匹配非数字 (相当于 \D)|
|[\D]|	not digits (≡ \D)|匹配非数字 (相当于 \D)|
|[^\D]|	not not digits (≡ \d)|匹配数字 (相当于 \d)|
|[[:name:]]|	named ASCII class inside character class (≡ [:name:])|命名的“ASCII 类”包含在“字符类”中 (相当于 [:name:])|
|[^[:name:]]|	named ASCII class inside negated character class (≡ [:^name:])|命名的“ASCII 类”不包含在“字符类”中 (相当于 [:^name:])|
|[\p{Name}]|	named Unicode property inside character class (≡ \p{Name})|命名的“Unicode 类”包含在“字符类”中 (相当于 \p{Name})|
|[^\p{Name}]|	named Unicode property inside negated character class (≡ \P{Name})|命名的“Unicode 类”不包含在“字符类”中 (相当于 \P{Name})|


|expression|Perl character classes (all ASCII-only)|Perl字符类（全部为ASCII码）|
|\d|	digits (≡ [0-9])|匹配数字 (相当于 [0-9])|
|\D|	not digits (≡ [^0-9])|匹配非数字 (相当于 [^0-9])|
|\s|	whitespace (≡ [\t\n\f\r ])|匹配空白 (相当于 [\t\n\f\r ])|
|\S|	not whitespace (≡ [^\t\n\f\r ])|匹配非空白 (相当于 [^\t\n\f\r ])|
|\w|	word characters (≡ [0-9A-Za-z_])|匹配单词字符 (相当于 [0-9A-Za-z_])|
|\W|	not word characters (≡ [^0-9A-Za-z_])|匹配非单词字符 (相当于 [^0-9A-Za-z_])|
|\h|	horizontal space (NOT SUPPORTED)|水平空间(不支持)|
|\H|	not horizontal space (NOT SUPPORTED)|非水平空间(不支持)|
|\v|	vertical space (NOT SUPPORTED)|垂直空间(不支持)|
|\V|	not vertical space (NOT SUPPORTED)|非垂直空间(不支持)|


|expression|ASCII character classes|ASCII字符类|
|:-|:-|:-|
|[[:alnum:]]|	alphanumeric (≡ [0-9A-Za-z])|字母数字 (相当于 [0-9A-Za-z])|
|[[:alpha:]]|	alphabetic (≡ [A-Za-z])|字母 (相当于 [A-Za-z])|
|[[:ascii:]]|	ASCII (≡ [\x00-\x7F])|ASCII 字符集 (相当于 [\x00-\x7F])|
|[[:blank:]]|	blank (≡ [\t ])|空白占位符 (相当于 [\t ])|
|[[:cntrl:]]|	control (≡ [\x00-\x1F\x7F])|控制字符 (相当于 [\x00-\x1F\x7F])|
|[[:digit:]]|	digits (≡ [0-9])|数字 (相当于 [0-9])|
|[[:graph:]]|	graphical (≡ [!-~] ≡ [A-Za-z0-9!"#$%&'()*+,\-./:;<=>?@[\\\]^_`{|}~])|图形字符 (相当于 [!-~])|
|[[:lower:]]|	lower case (≡ [a-z])|小写字母 (相当于 [a-z])|
|[[:print:]]|	printable (≡ [ -~] ≡ [ [:graph:]])|可打印字符 (相当于 [ -~] 相当于 [ [:graph:]])|
|[[:punct:]]|	punctuation (≡ [!-/:-@[-`{-~])|标点符号 (相当于 [!-/:-@[-反引号{-~])|
|[[:space:]]|	whitespace (≡ [\t\n\v\f\r ])|空白字符(相当于 [\t\n\v\f\r ])|
|[[:upper:]]|	upper case (≡ [A-Z])|大写字母(相当于 [A-Z])|
|[[:word:]]|	word characters (≡ [0-9A-Za-z_])|单词字符(相当于 [0-9A-Za-z_])|
|[[:xdigit:]]|	hex digit (≡ [0-9A-Fa-f])|16 进制字符集(相当于 [0-9A-Fa-f])|

|expression|Unicode character class names--general category|Unicode字符类名|
|:-|:-|:-|
|C|	other|其他|
|Cc|	control|控制字符|
|Cf|	format|格式|
|Cn|	unassigned code points (NOT SUPPORTED)|未分配的代码点（不支持）|
|Co|	private use|私有使用|
|Cs|	surrogate|代理|
|L|	letter|字母|
|LC|	cased letter (NOT SUPPORTED)|包装字母(不支持)|
|L&|	cased letter (NOT SUPPORTED)|包装字母(不支持)|
|Ll|	lowercase letter|小写字母|
|Lm|	modifier letter|修饰字母|
|Lo|	other letter|其他字母|
|Lt|	titlecase letter|首字母大写|
|Lu|	uppercase letter|大写字母|
|M|	mark|标记|
|Mc|	spacing mark|间隔标记|
|Me|	enclosing mark|附上标记|
|Mn|	non-spacing mark|非间隔标记|
|N|	number|数字|
|Nd|	decimal number|十进制数字|
|Nl|	letter number|字母数字|
|No|	other number|其他数字|
|P|	punctuation|标点符号|
|Pc|	connector punctuation|连接标点符号|
|Pd|	dash punctuation|破折号标点符号|
|Pe|	close punctuation|关闭的标点符号|
|Pf|	final punctuation|最后的标点符号|
|Pi|	initial punctuation|最初的标点符号|
|Po|	other punctuation|其他标点符号|
|Ps|	open punctuation|开放的标点符号|
|S|	symbol|符号|
|Sc|	currency symbol|货币符号|
|Sk|	modifier symbol|修饰符号|
|Sm|	math symbol|数学符号|
|So|	other symbol|其他符号|
|Z|	separator|分隔符|
|Zl|	line separator|行分隔符|
|Zp|	paragraph separator|段落分隔符|
|Zs|	space separator|空白分隔符|

Unicode character class names--scripts
Adlam
Ahom
Anatolian_Hieroglyphs
Arabic
Armenian
Avestan
Balinese
Bamum
Bassa_Vah
Batak
Bengali
Bhaiksuki
Bopomofo
Brahmi
Braille
Buginese
Buhid
Canadian_Aboriginal
Carian
Caucasian_Albanian
Chakma
Cham
Cherokee
Common
Coptic
Cuneiform
Cypriot
Cyrillic
Deseret
Devanagari
Duployan
Egyptian_Hieroglyphs
Elbasan
Ethiopic
Georgian
Glagolitic
Gothic
Grantha
Greek
Gujarati
Gurmukhi
Han
Hangul
Hanunoo
Hatran
Hebrew
Hiragana
Imperial_Aramaic
Inherited
Inscriptional_Pahlavi
Inscriptional_Parthian
Javanese
Kaithi
Kannada
Katakana
Kayah_Li
Kharoshthi
Khmer
Khojki
Khudawadi
Lao
Latin
Lepcha
Limbu
Linear_A
Linear_B
Lisu
Lycian
Lydian
Mahajani
Malayalam
Mandaic
Manichaean
Marchen
Masaram_Gondi
Meetei_Mayek
Mende_Kikakui
Meroitic_Cursive
Meroitic_Hieroglyphs
Miao
Modi
Mongolian
Mro
Multani
Myanmar
Nabataean
New_Tai_Lue
Newa
Nko
Nushu
Ogham
Ol_Chiki
Old_Hungarian
Old_Italic
Old_North_Arabian
Old_Permic
Old_Persian
Old_South_Arabian
Old_Turkic
Oriya
Osage
Osmanya
Pahawh_Hmong
Palmyrene
Pau_Cin_Hau
Phags_Pa
Phoenician
Psalter_Pahlavi
Rejang
Runic
Samaritan
Saurashtra
Sharada
Shavian
Siddham
SignWriting
Sinhala
Sora_Sompeng
Soyombo
Sundanese
Syloti_Nagri
Syriac
Tagalog
Tagbanwa
Tai_Le
Tai_Tham
Tai_Viet
Takri
Tamil
Tangut
Telugu
Thaana
Thai
Tibetan
Tifinagh
Tirhuta
Ugaritic
Vai
Warang_Citi
Yi
Zanabazar_Square
Vim character classes
\i	identifier character (NOT SUPPORTED) VIM
\I	\i except digits (NOT SUPPORTED) VIM
\k	keyword character (NOT SUPPORTED) VIM
\K	\k except digits (NOT SUPPORTED) VIM
\f	file name character (NOT SUPPORTED) VIM
\F	\f except digits (NOT SUPPORTED) VIM
\p	printable character (NOT SUPPORTED) VIM
\P	\p except digits (NOT SUPPORTED) VIM
\s	whitespace character (≡ [ \t]) (NOT SUPPORTED) VIM
\S	non-white space character (≡ [^ \t]) (NOT SUPPORTED) VIM
\d	digits (≡ [0-9]) VIM
\D	not \d VIM
\x	hex digits (≡ [0-9A-Fa-f]) (NOT SUPPORTED) VIM
\X	not \x (NOT SUPPORTED) VIM
\o	octal digits (≡ [0-7]) (NOT SUPPORTED) VIM
\O	not \o (NOT SUPPORTED) VIM
\w	word character VIM
\W	not \w VIM
\h	head of word character (NOT SUPPORTED) VIM
\H	not \h (NOT SUPPORTED) VIM
\a	alphabetic (NOT SUPPORTED) VIM
\A	not \a (NOT SUPPORTED) VIM
\l	lowercase (NOT SUPPORTED) VIM
\L	not lowercase (NOT SUPPORTED) VIM
\u	uppercase (NOT SUPPORTED) VIM
\U	not uppercase (NOT SUPPORTED) VIM
\_x	\x plus newline, for any x (NOT SUPPORTED) VIM
\c	ignore case (NOT SUPPORTED) VIM
\C	match case (NOT SUPPORTED) VIM
\m	magic (NOT SUPPORTED) VIM
\M	nomagic (NOT SUPPORTED) VIM
\v	verymagic (NOT SUPPORTED) VIM
\V	verynomagic (NOT SUPPORTED) VIM
\Z	ignore differences in Unicode combining characters (NOT SUPPORTED) VIM
Magic
(?{code})	arbitrary Perl code (NOT SUPPORTED) PERL
(??{code})	postponed arbitrary Perl code (NOT SUPPORTED) PERL
(?n)	recursive call to regexp capturing group n (NOT SUPPORTED)
(?+n)	recursive call to relative group +n (NOT SUPPORTED)
(?-n)	recursive call to relative group -n (NOT SUPPORTED)
(?C)	PCRE callout (NOT SUPPORTED) PCRE
(?R)	recursive call to entire regexp (≡ (?0)) (NOT SUPPORTED)
(?&name)	recursive call to named group (NOT SUPPORTED)
(?P=name)	named backreference (NOT SUPPORTED)
(?P>name)	recursive call to named group (NOT SUPPORTED)
(?(cond)true|false)	conditional branch (NOT SUPPORTED)
(?(cond)true)	conditional branch (NOT SUPPORTED)
(*ACCEPT)	make regexps more like Prolog (NOT SUPPORTED)
(*COMMIT)	(NOT SUPPORTED)
(*F)	(NOT SUPPORTED)
(*FAIL)	(NOT SUPPORTED)
(*MARK)	(NOT SUPPORTED)
(*PRUNE)	(NOT SUPPORTED)
(*SKIP)	(NOT SUPPORTED)
(*THEN)	(NOT SUPPORTED)
(*ANY)	set newline convention (NOT SUPPORTED)
(*ANYCRLF)	(NOT SUPPORTED)
(*CR)	(NOT SUPPORTED)
(*CRLF)	(NOT SUPPORTED)
(*LF)	(NOT SUPPORTED)
(*BSR_ANYCRLF)	set \R convention (NOT SUPPORTED) PCRE
(*BSR_UNICODE)	(NOT SUPPORTED) PCRE
