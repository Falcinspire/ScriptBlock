// Generated from c:\Falcinspire\go\src\github.com\falcinspire\scriptblock\front\parser\ScriptBlockParser.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ScriptBlockParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		WS=1, COMMENT=2, DOC_START=3, DOC_END=4, DOC_LINE=5, NEWLINES=6, FUNCTION=7, 
		TEMPLATE=8, NATIVE=9, CONSTANT=10, RUN=11, DELAY=12, IMPORT=13, INTERNAL=14, 
		ARROW=15, OPAREN=16, CPAREN=17, OCURLY=18, CCURLY=19, OSQUARE=20, CSQUARE=21, 
		COMMA=22, EQUALS=23, SEMICOLON=24, COLON=25, POWER=26, MULTIPLY=27, DIVIDE=28, 
		INTEGER_DIVIDE=29, PLUS=30, SUBTRACT=31, DOT=32, GREATER_THAN=33, LESS_THAN=34, 
		POUND=35, IDENTIFIER=36, SIGN=37, DIGITS=38, STRING=39, RAISE=40;
	public static final int
		RULE_unit = 0, RULE_documentation = 1, RULE_parameterList = 2, RULE_argumentList = 3, 
		RULE_structureList = 4, RULE_nativeCall = 5, RULE_functionCall = 6, RULE_functionFrame = 7, 
		RULE_functionDefinition = 8, RULE_functionDefinitionShortcut = 9, RULE_tag = 10, 
		RULE_templateDefinition = 11, RULE_constantDefinition = 12, RULE_formatter = 13, 
		RULE_importLine = 14, RULE_topDefinition = 15, RULE_statement = 16, RULE_delayStructure = 17, 
		RULE_raise = 18, RULE_expression = 19, RULE_number = 20;
	public static final String[] ruleNames = {
		"unit", "documentation", "parameterList", "argumentList", "structureList", 
		"nativeCall", "functionCall", "functionFrame", "functionDefinition", "functionDefinitionShortcut", 
		"tag", "templateDefinition", "constantDefinition", "formatter", "importLine", 
		"topDefinition", "statement", "delayStructure", "raise", "expression", 
		"number"
	};

	private static final String[] _LITERAL_NAMES = {
		null, null, null, "'/*'", "'*/'", null, null, "'func'", "'script'", "'command'", 
		"'const'", "'run'", "'delay'", "'import'", "'internal'", "'->'", "'('", 
		"')'", "'{'", "'}'", "'['", "']'", "','", "'='", "';'", "':'", "'^'", 
		"'*'", "'/'", "'//'", "'+'", "'-'", "'.'", "'>'", "'<'", "'#'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, "WS", "COMMENT", "DOC_START", "DOC_END", "DOC_LINE", "NEWLINES", 
		"FUNCTION", "TEMPLATE", "NATIVE", "CONSTANT", "RUN", "DELAY", "IMPORT", 
		"INTERNAL", "ARROW", "OPAREN", "CPAREN", "OCURLY", "CCURLY", "OSQUARE", 
		"CSQUARE", "COMMA", "EQUALS", "SEMICOLON", "COLON", "POWER", "MULTIPLY", 
		"DIVIDE", "INTEGER_DIVIDE", "PLUS", "SUBTRACT", "DOT", "GREATER_THAN", 
		"LESS_THAN", "POUND", "IDENTIFIER", "SIGN", "DIGITS", "STRING", "RAISE"
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

	@Override
	public String getGrammarFileName() { return "ScriptBlockParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public ScriptBlockParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}
	public static class UnitContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(ScriptBlockParser.EOF, 0); }
		public List<ImportLineContext> importLine() {
			return getRuleContexts(ImportLineContext.class);
		}
		public ImportLineContext importLine(int i) {
			return getRuleContext(ImportLineContext.class,i);
		}
		public List<TerminalNode> NEWLINES() { return getTokens(ScriptBlockParser.NEWLINES); }
		public TerminalNode NEWLINES(int i) {
			return getToken(ScriptBlockParser.NEWLINES, i);
		}
		public List<TopDefinitionContext> topDefinition() {
			return getRuleContexts(TopDefinitionContext.class);
		}
		public TopDefinitionContext topDefinition(int i) {
			return getRuleContext(TopDefinitionContext.class,i);
		}
		public UnitContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unit; }
	}

	public final UnitContext unit() throws RecognitionException {
		UnitContext _localctx = new UnitContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_unit);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(47);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IMPORT) {
				{
				{
				setState(42);
				importLine();
				setState(43);
				match(NEWLINES);
				}
				}
				setState(49);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(53);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << DOC_START) | (1L << FUNCTION) | (1L << TEMPLATE) | (1L << CONSTANT) | (1L << INTERNAL) | (1L << POUND))) != 0)) {
				{
				{
				setState(50);
				topDefinition();
				}
				}
				setState(55);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(56);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DocumentationContext extends ParserRuleContext {
		public TerminalNode DOC_START() { return getToken(ScriptBlockParser.DOC_START, 0); }
		public List<TerminalNode> NEWLINES() { return getTokens(ScriptBlockParser.NEWLINES); }
		public TerminalNode NEWLINES(int i) {
			return getToken(ScriptBlockParser.NEWLINES, i);
		}
		public TerminalNode DOC_END() { return getToken(ScriptBlockParser.DOC_END, 0); }
		public List<TerminalNode> DOC_LINE() { return getTokens(ScriptBlockParser.DOC_LINE); }
		public TerminalNode DOC_LINE(int i) {
			return getToken(ScriptBlockParser.DOC_LINE, i);
		}
		public DocumentationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_documentation; }
	}

	public final DocumentationContext documentation() throws RecognitionException {
		DocumentationContext _localctx = new DocumentationContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_documentation);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(58);
			match(DOC_START);
			setState(59);
			match(NEWLINES);
			setState(63);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOC_LINE) {
				{
				{
				setState(60);
				match(DOC_LINE);
				}
				}
				setState(65);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(66);
			match(DOC_END);
			setState(67);
			match(NEWLINES);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParameterListContext extends ParserRuleContext {
		public TerminalNode OPAREN() { return getToken(ScriptBlockParser.OPAREN, 0); }
		public TerminalNode CPAREN() { return getToken(ScriptBlockParser.CPAREN, 0); }
		public List<TerminalNode> IDENTIFIER() { return getTokens(ScriptBlockParser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(ScriptBlockParser.IDENTIFIER, i);
		}
		public List<TerminalNode> COMMA() { return getTokens(ScriptBlockParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(ScriptBlockParser.COMMA, i);
		}
		public ParameterListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameterList; }
	}

	public final ParameterListContext parameterList() throws RecognitionException {
		ParameterListContext _localctx = new ParameterListContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_parameterList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(69);
			match(OPAREN);
			setState(78);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENTIFIER) {
				{
				setState(70);
				match(IDENTIFIER);
				setState(75);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(71);
					match(COMMA);
					setState(72);
					match(IDENTIFIER);
					}
					}
					setState(77);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(80);
			match(CPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ArgumentListContext extends ParserRuleContext {
		public TerminalNode OPAREN() { return getToken(ScriptBlockParser.OPAREN, 0); }
		public TerminalNode CPAREN() { return getToken(ScriptBlockParser.CPAREN, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(ScriptBlockParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(ScriptBlockParser.COMMA, i);
		}
		public ArgumentListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_argumentList; }
	}

	public final ArgumentListContext argumentList() throws RecognitionException {
		ArgumentListContext _localctx = new ArgumentListContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_argumentList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(82);
			match(OPAREN);
			setState(91);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << RUN) | (1L << OPAREN) | (1L << LESS_THAN) | (1L << IDENTIFIER) | (1L << SIGN) | (1L << DIGITS) | (1L << STRING))) != 0)) {
				{
				setState(83);
				expression(0);
				setState(88);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(84);
					match(COMMA);
					setState(85);
					expression(0);
					}
					}
					setState(90);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(93);
			match(CPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StructureListContext extends ParserRuleContext {
		public TerminalNode OSQUARE() { return getToken(ScriptBlockParser.OSQUARE, 0); }
		public TerminalNode CSQUARE() { return getToken(ScriptBlockParser.CSQUARE, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(ScriptBlockParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(ScriptBlockParser.COMMA, i);
		}
		public StructureListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structureList; }
	}

	public final StructureListContext structureList() throws RecognitionException {
		StructureListContext _localctx = new StructureListContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_structureList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(95);
			match(OSQUARE);
			setState(104);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << RUN) | (1L << OPAREN) | (1L << LESS_THAN) | (1L << IDENTIFIER) | (1L << SIGN) | (1L << DIGITS) | (1L << STRING))) != 0)) {
				{
				setState(96);
				expression(0);
				setState(101);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(97);
					match(COMMA);
					setState(98);
					expression(0);
					}
					}
					setState(103);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(106);
			match(CSQUARE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NativeCallContext extends ParserRuleContext {
		public TerminalNode NATIVE() { return getToken(ScriptBlockParser.NATIVE, 0); }
		public ArgumentListContext argumentList() {
			return getRuleContext(ArgumentListContext.class,0);
		}
		public NativeCallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_nativeCall; }
	}

	public final NativeCallContext nativeCall() throws RecognitionException {
		NativeCallContext _localctx = new NativeCallContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_nativeCall);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(108);
			match(NATIVE);
			setState(109);
			argumentList();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FunctionCallContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(ScriptBlockParser.IDENTIFIER, 0); }
		public FunctionFrameContext functionFrame() {
			return getRuleContext(FunctionFrameContext.class,0);
		}
		public ArgumentListContext argumentList() {
			return getRuleContext(ArgumentListContext.class,0);
		}
		public FunctionCallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionCall; }
	}

	public final FunctionCallContext functionCall() throws RecognitionException {
		FunctionCallContext _localctx = new FunctionCallContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_functionCall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(111);
			match(IDENTIFIER);
			setState(117);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case OPAREN:
				{
				{
				setState(112);
				argumentList();
				setState(114);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==OCURLY) {
					{
					setState(113);
					functionFrame();
					}
				}

				}
				}
				break;
			case OCURLY:
				{
				setState(116);
				functionFrame();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FunctionFrameContext extends ParserRuleContext {
		public TerminalNode OCURLY() { return getToken(ScriptBlockParser.OCURLY, 0); }
		public TerminalNode NEWLINES() { return getToken(ScriptBlockParser.NEWLINES, 0); }
		public TerminalNode CCURLY() { return getToken(ScriptBlockParser.CCURLY, 0); }
		public ParameterListContext parameterList() {
			return getRuleContext(ParameterListContext.class,0);
		}
		public TerminalNode ARROW() { return getToken(ScriptBlockParser.ARROW, 0); }
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public FunctionFrameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionFrame; }
	}

	public final FunctionFrameContext functionFrame() throws RecognitionException {
		FunctionFrameContext _localctx = new FunctionFrameContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_functionFrame);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(119);
			match(OCURLY);
			setState(123);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==OPAREN) {
				{
				setState(120);
				parameterList();
				setState(121);
				match(ARROW);
				}
			}

			setState(125);
			match(NEWLINES);
			setState(129);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << NATIVE) | (1L << DELAY) | (1L << IDENTIFIER) | (1L << RAISE))) != 0)) {
				{
				{
				setState(126);
				statement();
				}
				}
				setState(131);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(132);
			match(CCURLY);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FunctionDefinitionContext extends ParserRuleContext {
		public TerminalNode FUNCTION() { return getToken(ScriptBlockParser.FUNCTION, 0); }
		public TerminalNode IDENTIFIER() { return getToken(ScriptBlockParser.IDENTIFIER, 0); }
		public FunctionFrameContext functionFrame() {
			return getRuleContext(FunctionFrameContext.class,0);
		}
		public DocumentationContext documentation() {
			return getRuleContext(DocumentationContext.class,0);
		}
		public TagContext tag() {
			return getRuleContext(TagContext.class,0);
		}
		public TerminalNode INTERNAL() { return getToken(ScriptBlockParser.INTERNAL, 0); }
		public FunctionDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionDefinition; }
	}

	public final FunctionDefinitionContext functionDefinition() throws RecognitionException {
		FunctionDefinitionContext _localctx = new FunctionDefinitionContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_functionDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(135);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DOC_START) {
				{
				setState(134);
				documentation();
				}
			}

			setState(138);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==POUND) {
				{
				setState(137);
				tag();
				}
			}

			setState(141);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==INTERNAL) {
				{
				setState(140);
				match(INTERNAL);
				}
			}

			setState(143);
			match(FUNCTION);
			setState(144);
			match(IDENTIFIER);
			setState(145);
			functionFrame();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FunctionDefinitionShortcutContext extends ParserRuleContext {
		public TerminalNode FUNCTION() { return getToken(ScriptBlockParser.FUNCTION, 0); }
		public TerminalNode IDENTIFIER() { return getToken(ScriptBlockParser.IDENTIFIER, 0); }
		public TerminalNode EQUALS() { return getToken(ScriptBlockParser.EQUALS, 0); }
		public FunctionCallContext functionCall() {
			return getRuleContext(FunctionCallContext.class,0);
		}
		public DocumentationContext documentation() {
			return getRuleContext(DocumentationContext.class,0);
		}
		public TagContext tag() {
			return getRuleContext(TagContext.class,0);
		}
		public TerminalNode INTERNAL() { return getToken(ScriptBlockParser.INTERNAL, 0); }
		public FunctionDefinitionShortcutContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionDefinitionShortcut; }
	}

	public final FunctionDefinitionShortcutContext functionDefinitionShortcut() throws RecognitionException {
		FunctionDefinitionShortcutContext _localctx = new FunctionDefinitionShortcutContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_functionDefinitionShortcut);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(148);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DOC_START) {
				{
				setState(147);
				documentation();
				}
			}

			setState(151);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==POUND) {
				{
				setState(150);
				tag();
				}
			}

			setState(154);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==INTERNAL) {
				{
				setState(153);
				match(INTERNAL);
				}
			}

			setState(156);
			match(FUNCTION);
			setState(157);
			match(IDENTIFIER);
			setState(158);
			match(EQUALS);
			setState(159);
			functionCall();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TagContext extends ParserRuleContext {
		public TerminalNode POUND() { return getToken(ScriptBlockParser.POUND, 0); }
		public List<TerminalNode> IDENTIFIER() { return getTokens(ScriptBlockParser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(ScriptBlockParser.IDENTIFIER, i);
		}
		public TerminalNode COLON() { return getToken(ScriptBlockParser.COLON, 0); }
		public TagContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tag; }
	}

	public final TagContext tag() throws RecognitionException {
		TagContext _localctx = new TagContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_tag);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(161);
			match(POUND);
			setState(162);
			match(IDENTIFIER);
			setState(163);
			match(COLON);
			setState(164);
			match(IDENTIFIER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TemplateDefinitionContext extends ParserRuleContext {
		public TerminalNode TEMPLATE() { return getToken(ScriptBlockParser.TEMPLATE, 0); }
		public TerminalNode IDENTIFIER() { return getToken(ScriptBlockParser.IDENTIFIER, 0); }
		public FunctionFrameContext functionFrame() {
			return getRuleContext(FunctionFrameContext.class,0);
		}
		public DocumentationContext documentation() {
			return getRuleContext(DocumentationContext.class,0);
		}
		public TerminalNode INTERNAL() { return getToken(ScriptBlockParser.INTERNAL, 0); }
		public TemplateDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_templateDefinition; }
	}

	public final TemplateDefinitionContext templateDefinition() throws RecognitionException {
		TemplateDefinitionContext _localctx = new TemplateDefinitionContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_templateDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(167);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DOC_START) {
				{
				setState(166);
				documentation();
				}
			}

			setState(170);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==INTERNAL) {
				{
				setState(169);
				match(INTERNAL);
				}
			}

			setState(172);
			match(TEMPLATE);
			setState(173);
			match(IDENTIFIER);
			setState(174);
			functionFrame();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ConstantDefinitionContext extends ParserRuleContext {
		public TerminalNode CONSTANT() { return getToken(ScriptBlockParser.CONSTANT, 0); }
		public TerminalNode IDENTIFIER() { return getToken(ScriptBlockParser.IDENTIFIER, 0); }
		public TerminalNode EQUALS() { return getToken(ScriptBlockParser.EQUALS, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public DocumentationContext documentation() {
			return getRuleContext(DocumentationContext.class,0);
		}
		public TerminalNode INTERNAL() { return getToken(ScriptBlockParser.INTERNAL, 0); }
		public ConstantDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_constantDefinition; }
	}

	public final ConstantDefinitionContext constantDefinition() throws RecognitionException {
		ConstantDefinitionContext _localctx = new ConstantDefinitionContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_constantDefinition);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(177);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DOC_START) {
				{
				setState(176);
				documentation();
				}
			}

			setState(180);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==INTERNAL) {
				{
				setState(179);
				match(INTERNAL);
				}
			}

			setState(182);
			match(CONSTANT);
			setState(183);
			match(IDENTIFIER);
			setState(184);
			match(EQUALS);
			setState(185);
			expression(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FormatterContext extends ParserRuleContext {
		public TerminalNode LESS_THAN() { return getToken(ScriptBlockParser.LESS_THAN, 0); }
		public TerminalNode IDENTIFIER() { return getToken(ScriptBlockParser.IDENTIFIER, 0); }
		public TerminalNode GREATER_THAN() { return getToken(ScriptBlockParser.GREATER_THAN, 0); }
		public ArgumentListContext argumentList() {
			return getRuleContext(ArgumentListContext.class,0);
		}
		public FormatterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_formatter; }
	}

	public final FormatterContext formatter() throws RecognitionException {
		FormatterContext _localctx = new FormatterContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_formatter);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(187);
			match(LESS_THAN);
			setState(188);
			match(IDENTIFIER);
			setState(189);
			match(GREATER_THAN);
			setState(190);
			argumentList();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ImportLineContext extends ParserRuleContext {
		public TerminalNode IMPORT() { return getToken(ScriptBlockParser.IMPORT, 0); }
		public List<TerminalNode> IDENTIFIER() { return getTokens(ScriptBlockParser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(ScriptBlockParser.IDENTIFIER, i);
		}
		public TerminalNode COLON() { return getToken(ScriptBlockParser.COLON, 0); }
		public ImportLineContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importLine; }
	}

	public final ImportLineContext importLine() throws RecognitionException {
		ImportLineContext _localctx = new ImportLineContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_importLine);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(192);
			match(IMPORT);
			setState(193);
			match(IDENTIFIER);
			setState(194);
			match(COLON);
			setState(195);
			match(IDENTIFIER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TopDefinitionContext extends ParserRuleContext {
		public TopDefinitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_topDefinition; }
	 
		public TopDefinitionContext() { }
		public void copyFrom(TopDefinitionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class FunctionDefinitionTopContext extends TopDefinitionContext {
		public FunctionDefinitionContext functionDefinition() {
			return getRuleContext(FunctionDefinitionContext.class,0);
		}
		public TerminalNode NEWLINES() { return getToken(ScriptBlockParser.NEWLINES, 0); }
		public FunctionDefinitionTopContext(TopDefinitionContext ctx) { copyFrom(ctx); }
	}
	public static class FunctionShortcutTopContext extends TopDefinitionContext {
		public FunctionDefinitionShortcutContext functionDefinitionShortcut() {
			return getRuleContext(FunctionDefinitionShortcutContext.class,0);
		}
		public TerminalNode NEWLINES() { return getToken(ScriptBlockParser.NEWLINES, 0); }
		public FunctionShortcutTopContext(TopDefinitionContext ctx) { copyFrom(ctx); }
	}
	public static class ConstantDefinitionTopContext extends TopDefinitionContext {
		public ConstantDefinitionContext constantDefinition() {
			return getRuleContext(ConstantDefinitionContext.class,0);
		}
		public TerminalNode NEWLINES() { return getToken(ScriptBlockParser.NEWLINES, 0); }
		public ConstantDefinitionTopContext(TopDefinitionContext ctx) { copyFrom(ctx); }
	}
	public static class TemplateDefinitionTopContext extends TopDefinitionContext {
		public TemplateDefinitionContext templateDefinition() {
			return getRuleContext(TemplateDefinitionContext.class,0);
		}
		public TerminalNode NEWLINES() { return getToken(ScriptBlockParser.NEWLINES, 0); }
		public TemplateDefinitionTopContext(TopDefinitionContext ctx) { copyFrom(ctx); }
	}

	public final TopDefinitionContext topDefinition() throws RecognitionException {
		TopDefinitionContext _localctx = new TopDefinitionContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_topDefinition);
		try {
			setState(209);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
			case 1:
				_localctx = new ConstantDefinitionTopContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(197);
				constantDefinition();
				setState(198);
				match(NEWLINES);
				}
				break;
			case 2:
				_localctx = new FunctionDefinitionTopContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(200);
				functionDefinition();
				setState(201);
				match(NEWLINES);
				}
				break;
			case 3:
				_localctx = new TemplateDefinitionTopContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(203);
				templateDefinition();
				setState(204);
				match(NEWLINES);
				}
				break;
			case 4:
				_localctx = new FunctionShortcutTopContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(206);
				functionDefinitionShortcut();
				setState(207);
				match(NEWLINES);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatementContext extends ParserRuleContext {
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	 
		public StatementContext() { }
		public void copyFrom(StatementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class RaiseStatementContext extends StatementContext {
		public RaiseContext raise() {
			return getRuleContext(RaiseContext.class,0);
		}
		public TerminalNode NEWLINES() { return getToken(ScriptBlockParser.NEWLINES, 0); }
		public RaiseStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	public static class FunctionCallStatementContext extends StatementContext {
		public FunctionCallContext functionCall() {
			return getRuleContext(FunctionCallContext.class,0);
		}
		public TerminalNode NEWLINES() { return getToken(ScriptBlockParser.NEWLINES, 0); }
		public FunctionCallStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	public static class DelayStructureStatementContext extends StatementContext {
		public DelayStructureContext delayStructure() {
			return getRuleContext(DelayStructureContext.class,0);
		}
		public TerminalNode NEWLINES() { return getToken(ScriptBlockParser.NEWLINES, 0); }
		public DelayStructureStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}
	public static class NativeCallStatementContext extends StatementContext {
		public NativeCallContext nativeCall() {
			return getRuleContext(NativeCallContext.class,0);
		}
		public TerminalNode NEWLINES() { return getToken(ScriptBlockParser.NEWLINES, 0); }
		public NativeCallStatementContext(StatementContext ctx) { copyFrom(ctx); }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_statement);
		try {
			setState(223);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IDENTIFIER:
				_localctx = new FunctionCallStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				{
				setState(211);
				functionCall();
				setState(212);
				match(NEWLINES);
				}
				}
				break;
			case NATIVE:
				_localctx = new NativeCallStatementContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				{
				setState(214);
				nativeCall();
				setState(215);
				match(NEWLINES);
				}
				}
				break;
			case DELAY:
				_localctx = new DelayStructureStatementContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				{
				setState(217);
				delayStructure();
				setState(218);
				match(NEWLINES);
				}
				}
				break;
			case RAISE:
				_localctx = new RaiseStatementContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				{
				setState(220);
				raise();
				setState(221);
				match(NEWLINES);
				}
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DelayStructureContext extends ParserRuleContext {
		public TerminalNode DELAY() { return getToken(ScriptBlockParser.DELAY, 0); }
		public StructureListContext structureList() {
			return getRuleContext(StructureListContext.class,0);
		}
		public FunctionFrameContext functionFrame() {
			return getRuleContext(FunctionFrameContext.class,0);
		}
		public DelayStructureContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_delayStructure; }
	}

	public final DelayStructureContext delayStructure() throws RecognitionException {
		DelayStructureContext _localctx = new DelayStructureContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_delayStructure);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(225);
			match(DELAY);
			setState(226);
			structureList();
			setState(227);
			functionFrame();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RaiseContext extends ParserRuleContext {
		public TerminalNode RAISE() { return getToken(ScriptBlockParser.RAISE, 0); }
		public TagContext tag() {
			return getRuleContext(TagContext.class,0);
		}
		public RaiseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_raise; }
	}

	public final RaiseContext raise() throws RecognitionException {
		RaiseContext _localctx = new RaiseContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_raise);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(229);
			match(RAISE);
			setState(230);
			tag();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionContext extends ParserRuleContext {
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	 
		public ExpressionContext() { }
		public void copyFrom(ExpressionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class StringExprContext extends ExpressionContext {
		public TerminalNode STRING() { return getToken(ScriptBlockParser.STRING, 0); }
		public StringExprContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	public static class DivideExprContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode DIVIDE() { return getToken(ScriptBlockParser.DIVIDE, 0); }
		public DivideExprContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	public static class IntegerDivideExprContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode INTEGER_DIVIDE() { return getToken(ScriptBlockParser.INTEGER_DIVIDE, 0); }
		public IntegerDivideExprContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	public static class SubtractExprContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode SUBTRACT() { return getToken(ScriptBlockParser.SUBTRACT, 0); }
		public SubtractExprContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	public static class PowerExprContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode POWER() { return getToken(ScriptBlockParser.POWER, 0); }
		public PowerExprContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	public static class AddExprContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode PLUS() { return getToken(ScriptBlockParser.PLUS, 0); }
		public AddExprContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	public static class NumberExprContext extends ExpressionContext {
		public NumberContext number() {
			return getRuleContext(NumberContext.class,0);
		}
		public NumberExprContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	public static class MultiplyExprContext extends ExpressionContext {
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public TerminalNode MULTIPLY() { return getToken(ScriptBlockParser.MULTIPLY, 0); }
		public MultiplyExprContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	public static class CallExprContext extends ExpressionContext {
		public TerminalNode RUN() { return getToken(ScriptBlockParser.RUN, 0); }
		public TerminalNode IDENTIFIER() { return getToken(ScriptBlockParser.IDENTIFIER, 0); }
		public ArgumentListContext argumentList() {
			return getRuleContext(ArgumentListContext.class,0);
		}
		public CallExprContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	public static class FormatterExprContext extends ExpressionContext {
		public FormatterContext formatter() {
			return getRuleContext(FormatterContext.class,0);
		}
		public FormatterExprContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	public static class IdentifierExprContext extends ExpressionContext {
		public TerminalNode IDENTIFIER() { return getToken(ScriptBlockParser.IDENTIFIER, 0); }
		public IdentifierExprContext(ExpressionContext ctx) { copyFrom(ctx); }
	}
	public static class ParenthExprContext extends ExpressionContext {
		public TerminalNode OPAREN() { return getToken(ScriptBlockParser.OPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode CPAREN() { return getToken(ScriptBlockParser.CPAREN, 0); }
		public ParenthExprContext(ExpressionContext ctx) { copyFrom(ctx); }
	}

	public final ExpressionContext expression() throws RecognitionException {
		return expression(0);
	}

	private ExpressionContext expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionContext _localctx = new ExpressionContext(_ctx, _parentState);
		ExpressionContext _prevctx = _localctx;
		int _startState = 38;
		enterRecursionRule(_localctx, 38, RULE_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(244);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SIGN:
			case DIGITS:
				{
				_localctx = new NumberExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(233);
				number();
				}
				break;
			case STRING:
				{
				_localctx = new StringExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(234);
				match(STRING);
				}
				break;
			case IDENTIFIER:
				{
				_localctx = new IdentifierExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(235);
				match(IDENTIFIER);
				}
				break;
			case OPAREN:
				{
				_localctx = new ParenthExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(236);
				match(OPAREN);
				setState(237);
				expression(0);
				setState(238);
				match(CPAREN);
				}
				break;
			case LESS_THAN:
				{
				_localctx = new FormatterExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(240);
				formatter();
				}
				break;
			case RUN:
				{
				_localctx = new CallExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(241);
				match(RUN);
				setState(242);
				match(IDENTIFIER);
				setState(243);
				argumentList();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(266);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,27,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(264);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,26,_ctx) ) {
					case 1:
						{
						_localctx = new PowerExprContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(246);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(247);
						match(POWER);
						setState(248);
						expression(7);
						}
						break;
					case 2:
						{
						_localctx = new MultiplyExprContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(249);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(250);
						match(MULTIPLY);
						setState(251);
						expression(7);
						}
						break;
					case 3:
						{
						_localctx = new DivideExprContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(252);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(253);
						match(DIVIDE);
						setState(254);
						expression(6);
						}
						break;
					case 4:
						{
						_localctx = new IntegerDivideExprContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(255);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(256);
						match(INTEGER_DIVIDE);
						setState(257);
						expression(5);
						}
						break;
					case 5:
						{
						_localctx = new AddExprContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(258);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(259);
						match(PLUS);
						setState(260);
						expression(4);
						}
						break;
					case 6:
						{
						_localctx = new SubtractExprContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(261);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(262);
						match(SUBTRACT);
						setState(263);
						expression(3);
						}
						break;
					}
					} 
				}
				setState(268);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,27,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class NumberContext extends ParserRuleContext {
		public List<TerminalNode> DIGITS() { return getTokens(ScriptBlockParser.DIGITS); }
		public TerminalNode DIGITS(int i) {
			return getToken(ScriptBlockParser.DIGITS, i);
		}
		public TerminalNode SIGN() { return getToken(ScriptBlockParser.SIGN, 0); }
		public TerminalNode DOT() { return getToken(ScriptBlockParser.DOT, 0); }
		public NumberContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_number; }
	}

	public final NumberContext number() throws RecognitionException {
		NumberContext _localctx = new NumberContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_number);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(270);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SIGN) {
				{
				setState(269);
				match(SIGN);
				}
			}

			setState(272);
			match(DIGITS);
			setState(275);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
			case 1:
				{
				setState(273);
				match(DOT);
				setState(274);
				match(DIGITS);
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 19:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 7);
		case 1:
			return precpred(_ctx, 6);
		case 2:
			return precpred(_ctx, 5);
		case 3:
			return precpred(_ctx, 4);
		case 4:
			return precpred(_ctx, 3);
		case 5:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3*\u0118\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\3\2\3\2\3\2\7\2\60\n\2\f\2\16"+
		"\2\63\13\2\3\2\7\2\66\n\2\f\2\16\29\13\2\3\2\3\2\3\3\3\3\3\3\7\3@\n\3"+
		"\f\3\16\3C\13\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\7\4L\n\4\f\4\16\4O\13\4\5"+
		"\4Q\n\4\3\4\3\4\3\5\3\5\3\5\3\5\7\5Y\n\5\f\5\16\5\\\13\5\5\5^\n\5\3\5"+
		"\3\5\3\6\3\6\3\6\3\6\7\6f\n\6\f\6\16\6i\13\6\5\6k\n\6\3\6\3\6\3\7\3\7"+
		"\3\7\3\b\3\b\3\b\5\bu\n\b\3\b\5\bx\n\b\3\t\3\t\3\t\3\t\5\t~\n\t\3\t\3"+
		"\t\7\t\u0082\n\t\f\t\16\t\u0085\13\t\3\t\3\t\3\n\5\n\u008a\n\n\3\n\5\n"+
		"\u008d\n\n\3\n\5\n\u0090\n\n\3\n\3\n\3\n\3\n\3\13\5\13\u0097\n\13\3\13"+
		"\5\13\u009a\n\13\3\13\5\13\u009d\n\13\3\13\3\13\3\13\3\13\3\13\3\f\3\f"+
		"\3\f\3\f\3\f\3\r\5\r\u00aa\n\r\3\r\5\r\u00ad\n\r\3\r\3\r\3\r\3\r\3\16"+
		"\5\16\u00b4\n\16\3\16\5\16\u00b7\n\16\3\16\3\16\3\16\3\16\3\16\3\17\3"+
		"\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3"+
		"\21\3\21\3\21\3\21\3\21\3\21\3\21\5\21\u00d4\n\21\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\5\22\u00e2\n\22\3\23\3\23\3\23"+
		"\3\23\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25"+
		"\3\25\3\25\5\25\u00f7\n\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25"+
		"\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\7\25\u010b\n\25\f\25\16"+
		"\25\u010e\13\25\3\26\5\26\u0111\n\26\3\26\3\26\3\26\5\26\u0116\n\26\3"+
		"\26\2\3(\27\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*\2\2\2\u012c"+
		"\2\61\3\2\2\2\4<\3\2\2\2\6G\3\2\2\2\bT\3\2\2\2\na\3\2\2\2\fn\3\2\2\2\16"+
		"q\3\2\2\2\20y\3\2\2\2\22\u0089\3\2\2\2\24\u0096\3\2\2\2\26\u00a3\3\2\2"+
		"\2\30\u00a9\3\2\2\2\32\u00b3\3\2\2\2\34\u00bd\3\2\2\2\36\u00c2\3\2\2\2"+
		" \u00d3\3\2\2\2\"\u00e1\3\2\2\2$\u00e3\3\2\2\2&\u00e7\3\2\2\2(\u00f6\3"+
		"\2\2\2*\u0110\3\2\2\2,-\5\36\20\2-.\7\b\2\2.\60\3\2\2\2/,\3\2\2\2\60\63"+
		"\3\2\2\2\61/\3\2\2\2\61\62\3\2\2\2\62\67\3\2\2\2\63\61\3\2\2\2\64\66\5"+
		" \21\2\65\64\3\2\2\2\669\3\2\2\2\67\65\3\2\2\2\678\3\2\2\28:\3\2\2\29"+
		"\67\3\2\2\2:;\7\2\2\3;\3\3\2\2\2<=\7\5\2\2=A\7\b\2\2>@\7\7\2\2?>\3\2\2"+
		"\2@C\3\2\2\2A?\3\2\2\2AB\3\2\2\2BD\3\2\2\2CA\3\2\2\2DE\7\6\2\2EF\7\b\2"+
		"\2F\5\3\2\2\2GP\7\22\2\2HM\7&\2\2IJ\7\30\2\2JL\7&\2\2KI\3\2\2\2LO\3\2"+
		"\2\2MK\3\2\2\2MN\3\2\2\2NQ\3\2\2\2OM\3\2\2\2PH\3\2\2\2PQ\3\2\2\2QR\3\2"+
		"\2\2RS\7\23\2\2S\7\3\2\2\2T]\7\22\2\2UZ\5(\25\2VW\7\30\2\2WY\5(\25\2X"+
		"V\3\2\2\2Y\\\3\2\2\2ZX\3\2\2\2Z[\3\2\2\2[^\3\2\2\2\\Z\3\2\2\2]U\3\2\2"+
		"\2]^\3\2\2\2^_\3\2\2\2_`\7\23\2\2`\t\3\2\2\2aj\7\26\2\2bg\5(\25\2cd\7"+
		"\30\2\2df\5(\25\2ec\3\2\2\2fi\3\2\2\2ge\3\2\2\2gh\3\2\2\2hk\3\2\2\2ig"+
		"\3\2\2\2jb\3\2\2\2jk\3\2\2\2kl\3\2\2\2lm\7\27\2\2m\13\3\2\2\2no\7\13\2"+
		"\2op\5\b\5\2p\r\3\2\2\2qw\7&\2\2rt\5\b\5\2su\5\20\t\2ts\3\2\2\2tu\3\2"+
		"\2\2ux\3\2\2\2vx\5\20\t\2wr\3\2\2\2wv\3\2\2\2x\17\3\2\2\2y}\7\24\2\2z"+
		"{\5\6\4\2{|\7\21\2\2|~\3\2\2\2}z\3\2\2\2}~\3\2\2\2~\177\3\2\2\2\177\u0083"+
		"\7\b\2\2\u0080\u0082\5\"\22\2\u0081\u0080\3\2\2\2\u0082\u0085\3\2\2\2"+
		"\u0083\u0081\3\2\2\2\u0083\u0084\3\2\2\2\u0084\u0086\3\2\2\2\u0085\u0083"+
		"\3\2\2\2\u0086\u0087\7\25\2\2\u0087\21\3\2\2\2\u0088\u008a\5\4\3\2\u0089"+
		"\u0088\3\2\2\2\u0089\u008a\3\2\2\2\u008a\u008c\3\2\2\2\u008b\u008d\5\26"+
		"\f\2\u008c\u008b\3\2\2\2\u008c\u008d\3\2\2\2\u008d\u008f\3\2\2\2\u008e"+
		"\u0090\7\20\2\2\u008f\u008e\3\2\2\2\u008f\u0090\3\2\2\2\u0090\u0091\3"+
		"\2\2\2\u0091\u0092\7\t\2\2\u0092\u0093\7&\2\2\u0093\u0094\5\20\t\2\u0094"+
		"\23\3\2\2\2\u0095\u0097\5\4\3\2\u0096\u0095\3\2\2\2\u0096\u0097\3\2\2"+
		"\2\u0097\u0099\3\2\2\2\u0098\u009a\5\26\f\2\u0099\u0098\3\2\2\2\u0099"+
		"\u009a\3\2\2\2\u009a\u009c\3\2\2\2\u009b\u009d\7\20\2\2\u009c\u009b\3"+
		"\2\2\2\u009c\u009d\3\2\2\2\u009d\u009e\3\2\2\2\u009e\u009f\7\t\2\2\u009f"+
		"\u00a0\7&\2\2\u00a0\u00a1\7\31\2\2\u00a1\u00a2\5\16\b\2\u00a2\25\3\2\2"+
		"\2\u00a3\u00a4\7%\2\2\u00a4\u00a5\7&\2\2\u00a5\u00a6\7\33\2\2\u00a6\u00a7"+
		"\7&\2\2\u00a7\27\3\2\2\2\u00a8\u00aa\5\4\3\2\u00a9\u00a8\3\2\2\2\u00a9"+
		"\u00aa\3\2\2\2\u00aa\u00ac\3\2\2\2\u00ab\u00ad\7\20\2\2\u00ac\u00ab\3"+
		"\2\2\2\u00ac\u00ad\3\2\2\2\u00ad\u00ae\3\2\2\2\u00ae\u00af\7\n\2\2\u00af"+
		"\u00b0\7&\2\2\u00b0\u00b1\5\20\t\2\u00b1\31\3\2\2\2\u00b2\u00b4\5\4\3"+
		"\2\u00b3\u00b2\3\2\2\2\u00b3\u00b4\3\2\2\2\u00b4\u00b6\3\2\2\2\u00b5\u00b7"+
		"\7\20\2\2\u00b6\u00b5\3\2\2\2\u00b6\u00b7\3\2\2\2\u00b7\u00b8\3\2\2\2"+
		"\u00b8\u00b9\7\f\2\2\u00b9\u00ba\7&\2\2\u00ba\u00bb\7\31\2\2\u00bb\u00bc"+
		"\5(\25\2\u00bc\33\3\2\2\2\u00bd\u00be\7$\2\2\u00be\u00bf\7&\2\2\u00bf"+
		"\u00c0\7#\2\2\u00c0\u00c1\5\b\5\2\u00c1\35\3\2\2\2\u00c2\u00c3\7\17\2"+
		"\2\u00c3\u00c4\7&\2\2\u00c4\u00c5\7\33\2\2\u00c5\u00c6\7&\2\2\u00c6\37"+
		"\3\2\2\2\u00c7\u00c8\5\32\16\2\u00c8\u00c9\7\b\2\2\u00c9\u00d4\3\2\2\2"+
		"\u00ca\u00cb\5\22\n\2\u00cb\u00cc\7\b\2\2\u00cc\u00d4\3\2\2\2\u00cd\u00ce"+
		"\5\30\r\2\u00ce\u00cf\7\b\2\2\u00cf\u00d4\3\2\2\2\u00d0\u00d1\5\24\13"+
		"\2\u00d1\u00d2\7\b\2\2\u00d2\u00d4\3\2\2\2\u00d3\u00c7\3\2\2\2\u00d3\u00ca"+
		"\3\2\2\2\u00d3\u00cd\3\2\2\2\u00d3\u00d0\3\2\2\2\u00d4!\3\2\2\2\u00d5"+
		"\u00d6\5\16\b\2\u00d6\u00d7\7\b\2\2\u00d7\u00e2\3\2\2\2\u00d8\u00d9\5"+
		"\f\7\2\u00d9\u00da\7\b\2\2\u00da\u00e2\3\2\2\2\u00db\u00dc\5$\23\2\u00dc"+
		"\u00dd\7\b\2\2\u00dd\u00e2\3\2\2\2\u00de\u00df\5&\24\2\u00df\u00e0\7\b"+
		"\2\2\u00e0\u00e2\3\2\2\2\u00e1\u00d5\3\2\2\2\u00e1\u00d8\3\2\2\2\u00e1"+
		"\u00db\3\2\2\2\u00e1\u00de\3\2\2\2\u00e2#\3\2\2\2\u00e3\u00e4\7\16\2\2"+
		"\u00e4\u00e5\5\n\6\2\u00e5\u00e6\5\20\t\2\u00e6%\3\2\2\2\u00e7\u00e8\7"+
		"*\2\2\u00e8\u00e9\5\26\f\2\u00e9\'\3\2\2\2\u00ea\u00eb\b\25\1\2\u00eb"+
		"\u00f7\5*\26\2\u00ec\u00f7\7)\2\2\u00ed\u00f7\7&\2\2\u00ee\u00ef\7\22"+
		"\2\2\u00ef\u00f0\5(\25\2\u00f0\u00f1\7\23\2\2\u00f1\u00f7\3\2\2\2\u00f2"+
		"\u00f7\5\34\17\2\u00f3\u00f4\7\r\2\2\u00f4\u00f5\7&\2\2\u00f5\u00f7\5"+
		"\b\5\2\u00f6\u00ea\3\2\2\2\u00f6\u00ec\3\2\2\2\u00f6\u00ed\3\2\2\2\u00f6"+
		"\u00ee\3\2\2\2\u00f6\u00f2\3\2\2\2\u00f6\u00f3\3\2\2\2\u00f7\u010c\3\2"+
		"\2\2\u00f8\u00f9\f\t\2\2\u00f9\u00fa\7\34\2\2\u00fa\u010b\5(\25\t\u00fb"+
		"\u00fc\f\b\2\2\u00fc\u00fd\7\35\2\2\u00fd\u010b\5(\25\t\u00fe\u00ff\f"+
		"\7\2\2\u00ff\u0100\7\36\2\2\u0100\u010b\5(\25\b\u0101\u0102\f\6\2\2\u0102"+
		"\u0103\7\37\2\2\u0103\u010b\5(\25\7\u0104\u0105\f\5\2\2\u0105\u0106\7"+
		" \2\2\u0106\u010b\5(\25\6\u0107\u0108\f\4\2\2\u0108\u0109\7!\2\2\u0109"+
		"\u010b\5(\25\5\u010a\u00f8\3\2\2\2\u010a\u00fb\3\2\2\2\u010a\u00fe\3\2"+
		"\2\2\u010a\u0101\3\2\2\2\u010a\u0104\3\2\2\2\u010a\u0107\3\2\2\2\u010b"+
		"\u010e\3\2\2\2\u010c\u010a\3\2\2\2\u010c\u010d\3\2\2\2\u010d)\3\2\2\2"+
		"\u010e\u010c\3\2\2\2\u010f\u0111\7\'\2\2\u0110\u010f\3\2\2\2\u0110\u0111"+
		"\3\2\2\2\u0111\u0112\3\2\2\2\u0112\u0115\7(\2\2\u0113\u0114\7\"\2\2\u0114"+
		"\u0116\7(\2\2\u0115\u0113\3\2\2\2\u0115\u0116\3\2\2\2\u0116+\3\2\2\2 "+
		"\61\67AMPZ]gjtw}\u0083\u0089\u008c\u008f\u0096\u0099\u009c\u00a9\u00ac"+
		"\u00b3\u00b6\u00d3\u00e1\u00f6\u010a\u010c\u0110\u0115";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}