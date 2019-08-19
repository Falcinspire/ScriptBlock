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
		POUND=35, IDENTIFIER=36, SIGN=37, DIGITS=38, STRING=39;
	public static final int
		RULE_unit = 0, RULE_documentation = 1, RULE_parameterList = 2, RULE_argumentList = 3, 
		RULE_structureList = 4, RULE_nativeCall = 5, RULE_functionCall = 6, RULE_functionFrame = 7, 
		RULE_functionDefinition = 8, RULE_functionDefinitionShortcut = 9, RULE_tag = 10, 
		RULE_templateDefinition = 11, RULE_constantDefinition = 12, RULE_formatter = 13, 
		RULE_importLine = 14, RULE_topDefinition = 15, RULE_statement = 16, RULE_delayStructure = 17, 
		RULE_expression = 18, RULE_number = 19;
	public static final String[] ruleNames = {
		"unit", "documentation", "parameterList", "argumentList", "structureList", 
		"nativeCall", "functionCall", "functionFrame", "functionDefinition", "functionDefinitionShortcut", 
		"tag", "templateDefinition", "constantDefinition", "formatter", "importLine", 
		"topDefinition", "statement", "delayStructure", "expression", "number"
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
			setState(45);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IMPORT) {
				{
				{
				setState(40);
				importLine();
				setState(41);
				match(NEWLINES);
				}
				}
				setState(47);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(51);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << DOC_START) | (1L << FUNCTION) | (1L << TEMPLATE) | (1L << CONSTANT) | (1L << INTERNAL) | (1L << POUND))) != 0)) {
				{
				{
				setState(48);
				topDefinition();
				}
				}
				setState(53);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(54);
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
			setState(56);
			match(DOC_START);
			setState(57);
			match(NEWLINES);
			setState(61);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOC_LINE) {
				{
				{
				setState(58);
				match(DOC_LINE);
				}
				}
				setState(63);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(64);
			match(DOC_END);
			setState(65);
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
			setState(67);
			match(OPAREN);
			setState(76);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENTIFIER) {
				{
				setState(68);
				match(IDENTIFIER);
				setState(73);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(69);
					match(COMMA);
					setState(70);
					match(IDENTIFIER);
					}
					}
					setState(75);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(78);
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
			setState(80);
			match(OPAREN);
			setState(89);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << RUN) | (1L << OPAREN) | (1L << LESS_THAN) | (1L << IDENTIFIER) | (1L << SIGN) | (1L << DIGITS) | (1L << STRING))) != 0)) {
				{
				setState(81);
				expression(0);
				setState(86);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(82);
					match(COMMA);
					setState(83);
					expression(0);
					}
					}
					setState(88);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(91);
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
			setState(93);
			match(OSQUARE);
			setState(102);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << RUN) | (1L << OPAREN) | (1L << LESS_THAN) | (1L << IDENTIFIER) | (1L << SIGN) | (1L << DIGITS) | (1L << STRING))) != 0)) {
				{
				setState(94);
				expression(0);
				setState(99);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(95);
					match(COMMA);
					setState(96);
					expression(0);
					}
					}
					setState(101);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(104);
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
			setState(106);
			match(NATIVE);
			setState(107);
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
			setState(109);
			match(IDENTIFIER);
			setState(115);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case OPAREN:
				{
				{
				setState(110);
				argumentList();
				setState(112);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==OCURLY) {
					{
					setState(111);
					functionFrame();
					}
				}

				}
				}
				break;
			case OCURLY:
				{
				setState(114);
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
			setState(117);
			match(OCURLY);
			setState(121);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==OPAREN) {
				{
				setState(118);
				parameterList();
				setState(119);
				match(ARROW);
				}
			}

			setState(123);
			match(NEWLINES);
			setState(127);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << NATIVE) | (1L << DELAY) | (1L << IDENTIFIER))) != 0)) {
				{
				{
				setState(124);
				statement();
				}
				}
				setState(129);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(130);
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
			setState(133);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DOC_START) {
				{
				setState(132);
				documentation();
				}
			}

			setState(136);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==POUND) {
				{
				setState(135);
				tag();
				}
			}

			setState(139);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==INTERNAL) {
				{
				setState(138);
				match(INTERNAL);
				}
			}

			setState(141);
			match(FUNCTION);
			setState(142);
			match(IDENTIFIER);
			setState(143);
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
			setState(146);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DOC_START) {
				{
				setState(145);
				documentation();
				}
			}

			setState(149);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==POUND) {
				{
				setState(148);
				tag();
				}
			}

			setState(152);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==INTERNAL) {
				{
				setState(151);
				match(INTERNAL);
				}
			}

			setState(154);
			match(FUNCTION);
			setState(155);
			match(IDENTIFIER);
			setState(156);
			match(EQUALS);
			setState(157);
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
			setState(159);
			match(POUND);
			setState(160);
			match(IDENTIFIER);
			setState(161);
			match(COLON);
			setState(162);
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
			setState(165);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DOC_START) {
				{
				setState(164);
				documentation();
				}
			}

			setState(168);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==INTERNAL) {
				{
				setState(167);
				match(INTERNAL);
				}
			}

			setState(170);
			match(TEMPLATE);
			setState(171);
			match(IDENTIFIER);
			setState(172);
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
			setState(175);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DOC_START) {
				{
				setState(174);
				documentation();
				}
			}

			setState(178);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==INTERNAL) {
				{
				setState(177);
				match(INTERNAL);
				}
			}

			setState(180);
			match(CONSTANT);
			setState(181);
			match(IDENTIFIER);
			setState(182);
			match(EQUALS);
			setState(183);
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
			setState(185);
			match(LESS_THAN);
			setState(186);
			match(IDENTIFIER);
			setState(187);
			match(GREATER_THAN);
			setState(188);
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
			setState(190);
			match(IMPORT);
			setState(191);
			match(IDENTIFIER);
			setState(192);
			match(COLON);
			setState(193);
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
			setState(207);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
			case 1:
				_localctx = new ConstantDefinitionTopContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(195);
				constantDefinition();
				setState(196);
				match(NEWLINES);
				}
				break;
			case 2:
				_localctx = new FunctionDefinitionTopContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(198);
				functionDefinition();
				setState(199);
				match(NEWLINES);
				}
				break;
			case 3:
				_localctx = new TemplateDefinitionTopContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(201);
				templateDefinition();
				setState(202);
				match(NEWLINES);
				}
				break;
			case 4:
				_localctx = new FunctionShortcutTopContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(204);
				functionDefinitionShortcut();
				setState(205);
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
			setState(218);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IDENTIFIER:
				_localctx = new FunctionCallStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				{
				setState(209);
				functionCall();
				setState(210);
				match(NEWLINES);
				}
				}
				break;
			case NATIVE:
				_localctx = new NativeCallStatementContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				{
				setState(212);
				nativeCall();
				setState(213);
				match(NEWLINES);
				}
				}
				break;
			case DELAY:
				_localctx = new DelayStructureStatementContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				{
				setState(215);
				delayStructure();
				setState(216);
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
			setState(220);
			match(DELAY);
			setState(221);
			structureList();
			setState(222);
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
		int _startState = 36;
		enterRecursionRule(_localctx, 36, RULE_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(236);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case SIGN:
			case DIGITS:
				{
				_localctx = new NumberExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(225);
				number();
				}
				break;
			case STRING:
				{
				_localctx = new StringExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(226);
				match(STRING);
				}
				break;
			case IDENTIFIER:
				{
				_localctx = new IdentifierExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(227);
				match(IDENTIFIER);
				}
				break;
			case OPAREN:
				{
				_localctx = new ParenthExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(228);
				match(OPAREN);
				setState(229);
				expression(0);
				setState(230);
				match(CPAREN);
				}
				break;
			case LESS_THAN:
				{
				_localctx = new FormatterExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(232);
				formatter();
				}
				break;
			case RUN:
				{
				_localctx = new CallExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(233);
				match(RUN);
				setState(234);
				match(IDENTIFIER);
				setState(235);
				argumentList();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(258);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,27,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(256);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,26,_ctx) ) {
					case 1:
						{
						_localctx = new PowerExprContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(238);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(239);
						match(POWER);
						setState(240);
						expression(7);
						}
						break;
					case 2:
						{
						_localctx = new MultiplyExprContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(241);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(242);
						match(MULTIPLY);
						setState(243);
						expression(7);
						}
						break;
					case 3:
						{
						_localctx = new DivideExprContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(244);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(245);
						match(DIVIDE);
						setState(246);
						expression(6);
						}
						break;
					case 4:
						{
						_localctx = new IntegerDivideExprContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(247);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(248);
						match(INTEGER_DIVIDE);
						setState(249);
						expression(5);
						}
						break;
					case 5:
						{
						_localctx = new AddExprContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(250);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(251);
						match(PLUS);
						setState(252);
						expression(4);
						}
						break;
					case 6:
						{
						_localctx = new SubtractExprContext(new ExpressionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expression);
						setState(253);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(254);
						match(SUBTRACT);
						setState(255);
						expression(3);
						}
						break;
					}
					} 
				}
				setState(260);
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
		enterRule(_localctx, 38, RULE_number);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(262);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SIGN) {
				{
				setState(261);
				match(SIGN);
				}
			}

			setState(264);
			match(DIGITS);
			setState(267);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
			case 1:
				{
				setState(265);
				match(DOT);
				setState(266);
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
		case 18:
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3)\u0110\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\3\2\3\2\3\2\7\2.\n\2\f\2\16\2\61\13\2\3"+
		"\2\7\2\64\n\2\f\2\16\2\67\13\2\3\2\3\2\3\3\3\3\3\3\7\3>\n\3\f\3\16\3A"+
		"\13\3\3\3\3\3\3\3\3\4\3\4\3\4\3\4\7\4J\n\4\f\4\16\4M\13\4\5\4O\n\4\3\4"+
		"\3\4\3\5\3\5\3\5\3\5\7\5W\n\5\f\5\16\5Z\13\5\5\5\\\n\5\3\5\3\5\3\6\3\6"+
		"\3\6\3\6\7\6d\n\6\f\6\16\6g\13\6\5\6i\n\6\3\6\3\6\3\7\3\7\3\7\3\b\3\b"+
		"\3\b\5\bs\n\b\3\b\5\bv\n\b\3\t\3\t\3\t\3\t\5\t|\n\t\3\t\3\t\7\t\u0080"+
		"\n\t\f\t\16\t\u0083\13\t\3\t\3\t\3\n\5\n\u0088\n\n\3\n\5\n\u008b\n\n\3"+
		"\n\5\n\u008e\n\n\3\n\3\n\3\n\3\n\3\13\5\13\u0095\n\13\3\13\5\13\u0098"+
		"\n\13\3\13\5\13\u009b\n\13\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3"+
		"\f\3\r\5\r\u00a8\n\r\3\r\5\r\u00ab\n\r\3\r\3\r\3\r\3\r\3\16\5\16\u00b2"+
		"\n\16\3\16\5\16\u00b5\n\16\3\16\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17"+
		"\3\17\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\5\21\u00d2\n\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\5\22\u00dd\n\22\3\23\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\24\5\24\u00ef\n\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\7\24\u0103\n\24\f\24\16\24\u0106\13\24\3\25\5\25\u0109\n\25\3\25\3\25"+
		"\3\25\5\25\u010e\n\25\3\25\2\3&\26\2\4\6\b\n\f\16\20\22\24\26\30\32\34"+
		"\36 \"$&(\2\2\2\u0124\2/\3\2\2\2\4:\3\2\2\2\6E\3\2\2\2\bR\3\2\2\2\n_\3"+
		"\2\2\2\fl\3\2\2\2\16o\3\2\2\2\20w\3\2\2\2\22\u0087\3\2\2\2\24\u0094\3"+
		"\2\2\2\26\u00a1\3\2\2\2\30\u00a7\3\2\2\2\32\u00b1\3\2\2\2\34\u00bb\3\2"+
		"\2\2\36\u00c0\3\2\2\2 \u00d1\3\2\2\2\"\u00dc\3\2\2\2$\u00de\3\2\2\2&\u00ee"+
		"\3\2\2\2(\u0108\3\2\2\2*+\5\36\20\2+,\7\b\2\2,.\3\2\2\2-*\3\2\2\2.\61"+
		"\3\2\2\2/-\3\2\2\2/\60\3\2\2\2\60\65\3\2\2\2\61/\3\2\2\2\62\64\5 \21\2"+
		"\63\62\3\2\2\2\64\67\3\2\2\2\65\63\3\2\2\2\65\66\3\2\2\2\668\3\2\2\2\67"+
		"\65\3\2\2\289\7\2\2\39\3\3\2\2\2:;\7\5\2\2;?\7\b\2\2<>\7\7\2\2=<\3\2\2"+
		"\2>A\3\2\2\2?=\3\2\2\2?@\3\2\2\2@B\3\2\2\2A?\3\2\2\2BC\7\6\2\2CD\7\b\2"+
		"\2D\5\3\2\2\2EN\7\22\2\2FK\7&\2\2GH\7\30\2\2HJ\7&\2\2IG\3\2\2\2JM\3\2"+
		"\2\2KI\3\2\2\2KL\3\2\2\2LO\3\2\2\2MK\3\2\2\2NF\3\2\2\2NO\3\2\2\2OP\3\2"+
		"\2\2PQ\7\23\2\2Q\7\3\2\2\2R[\7\22\2\2SX\5&\24\2TU\7\30\2\2UW\5&\24\2V"+
		"T\3\2\2\2WZ\3\2\2\2XV\3\2\2\2XY\3\2\2\2Y\\\3\2\2\2ZX\3\2\2\2[S\3\2\2\2"+
		"[\\\3\2\2\2\\]\3\2\2\2]^\7\23\2\2^\t\3\2\2\2_h\7\26\2\2`e\5&\24\2ab\7"+
		"\30\2\2bd\5&\24\2ca\3\2\2\2dg\3\2\2\2ec\3\2\2\2ef\3\2\2\2fi\3\2\2\2ge"+
		"\3\2\2\2h`\3\2\2\2hi\3\2\2\2ij\3\2\2\2jk\7\27\2\2k\13\3\2\2\2lm\7\13\2"+
		"\2mn\5\b\5\2n\r\3\2\2\2ou\7&\2\2pr\5\b\5\2qs\5\20\t\2rq\3\2\2\2rs\3\2"+
		"\2\2sv\3\2\2\2tv\5\20\t\2up\3\2\2\2ut\3\2\2\2v\17\3\2\2\2w{\7\24\2\2x"+
		"y\5\6\4\2yz\7\21\2\2z|\3\2\2\2{x\3\2\2\2{|\3\2\2\2|}\3\2\2\2}\u0081\7"+
		"\b\2\2~\u0080\5\"\22\2\177~\3\2\2\2\u0080\u0083\3\2\2\2\u0081\177\3\2"+
		"\2\2\u0081\u0082\3\2\2\2\u0082\u0084\3\2\2\2\u0083\u0081\3\2\2\2\u0084"+
		"\u0085\7\25\2\2\u0085\21\3\2\2\2\u0086\u0088\5\4\3\2\u0087\u0086\3\2\2"+
		"\2\u0087\u0088\3\2\2\2\u0088\u008a\3\2\2\2\u0089\u008b\5\26\f\2\u008a"+
		"\u0089\3\2\2\2\u008a\u008b\3\2\2\2\u008b\u008d\3\2\2\2\u008c\u008e\7\20"+
		"\2\2\u008d\u008c\3\2\2\2\u008d\u008e\3\2\2\2\u008e\u008f\3\2\2\2\u008f"+
		"\u0090\7\t\2\2\u0090\u0091\7&\2\2\u0091\u0092\5\20\t\2\u0092\23\3\2\2"+
		"\2\u0093\u0095\5\4\3\2\u0094\u0093\3\2\2\2\u0094\u0095\3\2\2\2\u0095\u0097"+
		"\3\2\2\2\u0096\u0098\5\26\f\2\u0097\u0096\3\2\2\2\u0097\u0098\3\2\2\2"+
		"\u0098\u009a\3\2\2\2\u0099\u009b\7\20\2\2\u009a\u0099\3\2\2\2\u009a\u009b"+
		"\3\2\2\2\u009b\u009c\3\2\2\2\u009c\u009d\7\t\2\2\u009d\u009e\7&\2\2\u009e"+
		"\u009f\7\31\2\2\u009f\u00a0\5\16\b\2\u00a0\25\3\2\2\2\u00a1\u00a2\7%\2"+
		"\2\u00a2\u00a3\7&\2\2\u00a3\u00a4\7\33\2\2\u00a4\u00a5\7&\2\2\u00a5\27"+
		"\3\2\2\2\u00a6\u00a8\5\4\3\2\u00a7\u00a6\3\2\2\2\u00a7\u00a8\3\2\2\2\u00a8"+
		"\u00aa\3\2\2\2\u00a9\u00ab\7\20\2\2\u00aa\u00a9\3\2\2\2\u00aa\u00ab\3"+
		"\2\2\2\u00ab\u00ac\3\2\2\2\u00ac\u00ad\7\n\2\2\u00ad\u00ae\7&\2\2\u00ae"+
		"\u00af\5\20\t\2\u00af\31\3\2\2\2\u00b0\u00b2\5\4\3\2\u00b1\u00b0\3\2\2"+
		"\2\u00b1\u00b2\3\2\2\2\u00b2\u00b4\3\2\2\2\u00b3\u00b5\7\20\2\2\u00b4"+
		"\u00b3\3\2\2\2\u00b4\u00b5\3\2\2\2\u00b5\u00b6\3\2\2\2\u00b6\u00b7\7\f"+
		"\2\2\u00b7\u00b8\7&\2\2\u00b8\u00b9\7\31\2\2\u00b9\u00ba\5&\24\2\u00ba"+
		"\33\3\2\2\2\u00bb\u00bc\7$\2\2\u00bc\u00bd\7&\2\2\u00bd\u00be\7#\2\2\u00be"+
		"\u00bf\5\b\5\2\u00bf\35\3\2\2\2\u00c0\u00c1\7\17\2\2\u00c1\u00c2\7&\2"+
		"\2\u00c2\u00c3\7\33\2\2\u00c3\u00c4\7&\2\2\u00c4\37\3\2\2\2\u00c5\u00c6"+
		"\5\32\16\2\u00c6\u00c7\7\b\2\2\u00c7\u00d2\3\2\2\2\u00c8\u00c9\5\22\n"+
		"\2\u00c9\u00ca\7\b\2\2\u00ca\u00d2\3\2\2\2\u00cb\u00cc\5\30\r\2\u00cc"+
		"\u00cd\7\b\2\2\u00cd\u00d2\3\2\2\2\u00ce\u00cf\5\24\13\2\u00cf\u00d0\7"+
		"\b\2\2\u00d0\u00d2\3\2\2\2\u00d1\u00c5\3\2\2\2\u00d1\u00c8\3\2\2\2\u00d1"+
		"\u00cb\3\2\2\2\u00d1\u00ce\3\2\2\2\u00d2!\3\2\2\2\u00d3\u00d4\5\16\b\2"+
		"\u00d4\u00d5\7\b\2\2\u00d5\u00dd\3\2\2\2\u00d6\u00d7\5\f\7\2\u00d7\u00d8"+
		"\7\b\2\2\u00d8\u00dd\3\2\2\2\u00d9\u00da\5$\23\2\u00da\u00db\7\b\2\2\u00db"+
		"\u00dd\3\2\2\2\u00dc\u00d3\3\2\2\2\u00dc\u00d6\3\2\2\2\u00dc\u00d9\3\2"+
		"\2\2\u00dd#\3\2\2\2\u00de\u00df\7\16\2\2\u00df\u00e0\5\n\6\2\u00e0\u00e1"+
		"\5\20\t\2\u00e1%\3\2\2\2\u00e2\u00e3\b\24\1\2\u00e3\u00ef\5(\25\2\u00e4"+
		"\u00ef\7)\2\2\u00e5\u00ef\7&\2\2\u00e6\u00e7\7\22\2\2\u00e7\u00e8\5&\24"+
		"\2\u00e8\u00e9\7\23\2\2\u00e9\u00ef\3\2\2\2\u00ea\u00ef\5\34\17\2\u00eb"+
		"\u00ec\7\r\2\2\u00ec\u00ed\7&\2\2\u00ed\u00ef\5\b\5\2\u00ee\u00e2\3\2"+
		"\2\2\u00ee\u00e4\3\2\2\2\u00ee\u00e5\3\2\2\2\u00ee\u00e6\3\2\2\2\u00ee"+
		"\u00ea\3\2\2\2\u00ee\u00eb\3\2\2\2\u00ef\u0104\3\2\2\2\u00f0\u00f1\f\t"+
		"\2\2\u00f1\u00f2\7\34\2\2\u00f2\u0103\5&\24\t\u00f3\u00f4\f\b\2\2\u00f4"+
		"\u00f5\7\35\2\2\u00f5\u0103\5&\24\t\u00f6\u00f7\f\7\2\2\u00f7\u00f8\7"+
		"\36\2\2\u00f8\u0103\5&\24\b\u00f9\u00fa\f\6\2\2\u00fa\u00fb\7\37\2\2\u00fb"+
		"\u0103\5&\24\7\u00fc\u00fd\f\5\2\2\u00fd\u00fe\7 \2\2\u00fe\u0103\5&\24"+
		"\6\u00ff\u0100\f\4\2\2\u0100\u0101\7!\2\2\u0101\u0103\5&\24\5\u0102\u00f0"+
		"\3\2\2\2\u0102\u00f3\3\2\2\2\u0102\u00f6\3\2\2\2\u0102\u00f9\3\2\2\2\u0102"+
		"\u00fc\3\2\2\2\u0102\u00ff\3\2\2\2\u0103\u0106\3\2\2\2\u0104\u0102\3\2"+
		"\2\2\u0104\u0105\3\2\2\2\u0105\'\3\2\2\2\u0106\u0104\3\2\2\2\u0107\u0109"+
		"\7\'\2\2\u0108\u0107\3\2\2\2\u0108\u0109\3\2\2\2\u0109\u010a\3\2\2\2\u010a"+
		"\u010d\7(\2\2\u010b\u010c\7\"\2\2\u010c\u010e\7(\2\2\u010d\u010b\3\2\2"+
		"\2\u010d\u010e\3\2\2\2\u010e)\3\2\2\2 /\65?KNX[ehru{\u0081\u0087\u008a"+
		"\u008d\u0094\u0097\u009a\u00a7\u00aa\u00b1\u00b4\u00d1\u00dc\u00ee\u0102"+
		"\u0104\u0108\u010d";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}