// Generated from c:\Falcinspire\go\src\github.com\falcinspire\scriptblock\front\parser/ScriptBlockLexer.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ScriptBlockLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		WS=1, COMMENT=2, DOC_START=3, DOC_END=4, DOC_LINE=5, NEWLINES=6, FUNCTION=7, 
		TEMPLATE=8, NATIVE=9, CONSTANT=10, RUN=11, RAISE=12, DELAY=13, IMPORT=14, 
		INTERNAL=15, ARROW=16, OPAREN=17, CPAREN=18, OCURLY=19, CCURLY=20, OSQUARE=21, 
		CSQUARE=22, COMMA=23, EQUALS=24, SEMICOLON=25, COLON=26, POWER=27, MULTIPLY=28, 
		DIVIDE=29, INTEGER_DIVIDE=30, PLUS=31, SUBTRACT=32, DOT=33, GREATER_THAN=34, 
		LESS_THAN=35, POUND=36, IDENTIFIER=37, SIGN=38, DIGITS=39, STRING=40;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"WS", "COMMENT", "DOC_START", "DOC_END", "DOC_LINE", "NEWLINES", "FUNCTION", 
		"TEMPLATE", "NATIVE", "CONSTANT", "RUN", "RAISE", "DELAY", "IMPORT", "INTERNAL", 
		"ARROW", "OPAREN", "CPAREN", "OCURLY", "CCURLY", "OSQUARE", "CSQUARE", 
		"COMMA", "EQUALS", "SEMICOLON", "COLON", "POWER", "MULTIPLY", "DIVIDE", 
		"INTEGER_DIVIDE", "PLUS", "SUBTRACT", "DOT", "GREATER_THAN", "LESS_THAN", 
		"POUND", "IDENTIFIER", "SIGN", "DIGITS", "STRING"
	};

	private static final String[] _LITERAL_NAMES = {
		null, null, null, "'/*'", "'*/'", null, null, "'func'", "'script'", "'command'", 
		"'const'", "'run'", "'raise'", "'delay'", "'import'", "'internal'", "'->'", 
		"'('", "')'", "'{'", "'}'", "'['", "']'", "','", "'='", "';'", "':'", 
		"'^'", "'*'", "'/'", "'//'", "'+'", "'-'", "'.'", "'>'", "'<'", "'#'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, "WS", "COMMENT", "DOC_START", "DOC_END", "DOC_LINE", "NEWLINES", 
		"FUNCTION", "TEMPLATE", "NATIVE", "CONSTANT", "RUN", "RAISE", "DELAY", 
		"IMPORT", "INTERNAL", "ARROW", "OPAREN", "CPAREN", "OCURLY", "CCURLY", 
		"OSQUARE", "CSQUARE", "COMMA", "EQUALS", "SEMICOLON", "COLON", "POWER", 
		"MULTIPLY", "DIVIDE", "INTEGER_DIVIDE", "PLUS", "SUBTRACT", "DOT", "GREATER_THAN", 
		"LESS_THAN", "POUND", "IDENTIFIER", "SIGN", "DIGITS", "STRING"
	};
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public ScriptBlockLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "ScriptBlockLexer.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2*\u00fe\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\3\2\6\2U\n\2\r"+
		"\2\16\2V\3\2\3\2\3\3\3\3\3\3\3\3\7\3_\n\3\f\3\16\3b\13\3\3\3\6\3e\n\3"+
		"\r\3\16\3f\3\3\3\3\3\4\3\4\3\4\3\5\3\5\3\5\3\6\3\6\3\6\3\6\7\6u\n\6\f"+
		"\6\16\6x\13\6\3\6\6\6{\n\6\r\6\16\6|\3\7\6\7\u0080\n\7\r\7\16\7\u0081"+
		"\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3"+
		"\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17"+
		"\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\22"+
		"\3\22\3\23\3\23\3\24\3\24\3\25\3\25\3\26\3\26\3\27\3\27\3\30\3\30\3\31"+
		"\3\31\3\32\3\32\3\33\3\33\3\34\3\34\3\35\3\35\3\36\3\36\3\37\3\37\3\37"+
		"\3 \3 \3!\3!\3\"\3\"\3#\3#\3$\3$\3%\3%\3&\6&\u00eb\n&\r&\16&\u00ec\3\'"+
		"\3\'\3(\6(\u00f2\n(\r(\16(\u00f3\3)\3)\7)\u00f8\n)\f)\16)\u00fb\13)\3"+
		")\3)\5`v\u00f9\2*\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31"+
		"\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65"+
		"\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*\3\2\t\4\2\13\13\"\"\5\2\f\f"+
		"\17\17^^\4\2\f\f\17\17\5\2C\\aac|\4\2--//\3\2\62;\3\2))\2\u0106\2\3\3"+
		"\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2"+
		"\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3"+
		"\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2"+
		"%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61"+
		"\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2"+
		"\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I"+
		"\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\3T\3\2\2\2\5Z\3\2"+
		"\2\2\7j\3\2\2\2\tm\3\2\2\2\13p\3\2\2\2\r\177\3\2\2\2\17\u0083\3\2\2\2"+
		"\21\u0088\3\2\2\2\23\u008f\3\2\2\2\25\u0097\3\2\2\2\27\u009d\3\2\2\2\31"+
		"\u00a1\3\2\2\2\33\u00a7\3\2\2\2\35\u00ad\3\2\2\2\37\u00b4\3\2\2\2!\u00bd"+
		"\3\2\2\2#\u00c0\3\2\2\2%\u00c2\3\2\2\2\'\u00c4\3\2\2\2)\u00c6\3\2\2\2"+
		"+\u00c8\3\2\2\2-\u00ca\3\2\2\2/\u00cc\3\2\2\2\61\u00ce\3\2\2\2\63\u00d0"+
		"\3\2\2\2\65\u00d2\3\2\2\2\67\u00d4\3\2\2\29\u00d6\3\2\2\2;\u00d8\3\2\2"+
		"\2=\u00da\3\2\2\2?\u00dd\3\2\2\2A\u00df\3\2\2\2C\u00e1\3\2\2\2E\u00e3"+
		"\3\2\2\2G\u00e5\3\2\2\2I\u00e7\3\2\2\2K\u00ea\3\2\2\2M\u00ee\3\2\2\2O"+
		"\u00f1\3\2\2\2Q\u00f5\3\2\2\2SU\t\2\2\2TS\3\2\2\2UV\3\2\2\2VT\3\2\2\2"+
		"VW\3\2\2\2WX\3\2\2\2XY\b\2\2\2Y\4\3\2\2\2Z[\7%\2\2[\\\7%\2\2\\`\3\2\2"+
		"\2]_\n\3\2\2^]\3\2\2\2_b\3\2\2\2`a\3\2\2\2`^\3\2\2\2ad\3\2\2\2b`\3\2\2"+
		"\2ce\t\4\2\2dc\3\2\2\2ef\3\2\2\2fd\3\2\2\2fg\3\2\2\2gh\3\2\2\2hi\b\3\2"+
		"\2i\6\3\2\2\2jk\7\61\2\2kl\7,\2\2l\b\3\2\2\2mn\7,\2\2no\7\61\2\2o\n\3"+
		"\2\2\2pq\7,\2\2qr\7,\2\2rv\3\2\2\2su\n\4\2\2ts\3\2\2\2ux\3\2\2\2vw\3\2"+
		"\2\2vt\3\2\2\2wz\3\2\2\2xv\3\2\2\2y{\t\4\2\2zy\3\2\2\2{|\3\2\2\2|z\3\2"+
		"\2\2|}\3\2\2\2}\f\3\2\2\2~\u0080\t\4\2\2\177~\3\2\2\2\u0080\u0081\3\2"+
		"\2\2\u0081\177\3\2\2\2\u0081\u0082\3\2\2\2\u0082\16\3\2\2\2\u0083\u0084"+
		"\7h\2\2\u0084\u0085\7w\2\2\u0085\u0086\7p\2\2\u0086\u0087\7e\2\2\u0087"+
		"\20\3\2\2\2\u0088\u0089\7u\2\2\u0089\u008a\7e\2\2\u008a\u008b\7t\2\2\u008b"+
		"\u008c\7k\2\2\u008c\u008d\7r\2\2\u008d\u008e\7v\2\2\u008e\22\3\2\2\2\u008f"+
		"\u0090\7e\2\2\u0090\u0091\7q\2\2\u0091\u0092\7o\2\2\u0092\u0093\7o\2\2"+
		"\u0093\u0094\7c\2\2\u0094\u0095\7p\2\2\u0095\u0096\7f\2\2\u0096\24\3\2"+
		"\2\2\u0097\u0098\7e\2\2\u0098\u0099\7q\2\2\u0099\u009a\7p\2\2\u009a\u009b"+
		"\7u\2\2\u009b\u009c\7v\2\2\u009c\26\3\2\2\2\u009d\u009e\7t\2\2\u009e\u009f"+
		"\7w\2\2\u009f\u00a0\7p\2\2\u00a0\30\3\2\2\2\u00a1\u00a2\7t\2\2\u00a2\u00a3"+
		"\7c\2\2\u00a3\u00a4\7k\2\2\u00a4\u00a5\7u\2\2\u00a5\u00a6\7g\2\2\u00a6"+
		"\32\3\2\2\2\u00a7\u00a8\7f\2\2\u00a8\u00a9\7g\2\2\u00a9\u00aa\7n\2\2\u00aa"+
		"\u00ab\7c\2\2\u00ab\u00ac\7{\2\2\u00ac\34\3\2\2\2\u00ad\u00ae\7k\2\2\u00ae"+
		"\u00af\7o\2\2\u00af\u00b0\7r\2\2\u00b0\u00b1\7q\2\2\u00b1\u00b2\7t\2\2"+
		"\u00b2\u00b3\7v\2\2\u00b3\36\3\2\2\2\u00b4\u00b5\7k\2\2\u00b5\u00b6\7"+
		"p\2\2\u00b6\u00b7\7v\2\2\u00b7\u00b8\7g\2\2\u00b8\u00b9\7t\2\2\u00b9\u00ba"+
		"\7p\2\2\u00ba\u00bb\7c\2\2\u00bb\u00bc\7n\2\2\u00bc \3\2\2\2\u00bd\u00be"+
		"\7/\2\2\u00be\u00bf\7@\2\2\u00bf\"\3\2\2\2\u00c0\u00c1\7*\2\2\u00c1$\3"+
		"\2\2\2\u00c2\u00c3\7+\2\2\u00c3&\3\2\2\2\u00c4\u00c5\7}\2\2\u00c5(\3\2"+
		"\2\2\u00c6\u00c7\7\177\2\2\u00c7*\3\2\2\2\u00c8\u00c9\7]\2\2\u00c9,\3"+
		"\2\2\2\u00ca\u00cb\7_\2\2\u00cb.\3\2\2\2\u00cc\u00cd\7.\2\2\u00cd\60\3"+
		"\2\2\2\u00ce\u00cf\7?\2\2\u00cf\62\3\2\2\2\u00d0\u00d1\7=\2\2\u00d1\64"+
		"\3\2\2\2\u00d2\u00d3\7<\2\2\u00d3\66\3\2\2\2\u00d4\u00d5\7`\2\2\u00d5"+
		"8\3\2\2\2\u00d6\u00d7\7,\2\2\u00d7:\3\2\2\2\u00d8\u00d9\7\61\2\2\u00d9"+
		"<\3\2\2\2\u00da\u00db\7\61\2\2\u00db\u00dc\7\61\2\2\u00dc>\3\2\2\2\u00dd"+
		"\u00de\7-\2\2\u00de@\3\2\2\2\u00df\u00e0\7/\2\2\u00e0B\3\2\2\2\u00e1\u00e2"+
		"\7\60\2\2\u00e2D\3\2\2\2\u00e3\u00e4\7@\2\2\u00e4F\3\2\2\2\u00e5\u00e6"+
		"\7>\2\2\u00e6H\3\2\2\2\u00e7\u00e8\7%\2\2\u00e8J\3\2\2\2\u00e9\u00eb\t"+
		"\5\2\2\u00ea\u00e9\3\2\2\2\u00eb\u00ec\3\2\2\2\u00ec\u00ea\3\2\2\2\u00ec"+
		"\u00ed\3\2\2\2\u00edL\3\2\2\2\u00ee\u00ef\t\6\2\2\u00efN\3\2\2\2\u00f0"+
		"\u00f2\t\7\2\2\u00f1\u00f0\3\2\2\2\u00f2\u00f3\3\2\2\2\u00f3\u00f1\3\2"+
		"\2\2\u00f3\u00f4\3\2\2\2\u00f4P\3\2\2\2\u00f5\u00f9\7)\2\2\u00f6\u00f8"+
		"\n\b\2\2\u00f7\u00f6\3\2\2\2\u00f8\u00fb\3\2\2\2\u00f9\u00fa\3\2\2\2\u00f9"+
		"\u00f7\3\2\2\2\u00fa\u00fc\3\2\2\2\u00fb\u00f9\3\2\2\2\u00fc\u00fd\7)"+
		"\2\2\u00fdR\3\2\2\2\f\2V`fv|\u0081\u00ec\u00f3\u00f9\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}