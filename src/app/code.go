package app

import (
	"bytes"
	"context"
	"io"
	"log"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/canpacis/pacis/html"
	"github.com/canpacis/pacis/lucide"
	"github.com/canpacis/pacis/x"
)

func gettokclass(typ chroma.TokenType) string {
	switch typ {
	case chroma.PreWrapper:
		return "whitespace-pre-wrap"

	case chroma.Line:
		return "flex"

	case chroma.LineTable:
		return "border-spacing-0 p-0 m-0 border-0"
	case chroma.LineTableTD:
		return "align-top p-0 m-0 border-0"

	case chroma.LineLink:
		return "outline-0 decoration-0 text-inherit"

	case chroma.LineNumbers,
		chroma.LineNumbersTable:
		return "whitespace-pre-wrap select-none mr-4 py-2 text-neutral-400 dark:text-neutral-600"

	case chroma.LineHighlight:
		return "bg-sky-600/40"

	case chroma.Error:
		return "text-red-500 dark:text-red-400 bg-red-950 dark:bg-red-800"

	case chroma.Keyword,
		chroma.KeywordConstant,
		chroma.KeywordDeclaration,
		chroma.KeywordPseudo,
		chroma.KeywordReserved,
		chroma.KeywordType:
		return "text-rose-600 dark:text-rose-400"

	case chroma.KeywordNamespace:
		return "text-indigo-600 dark:text-indigo-400"

	case chroma.NameAttribute,
		chroma.NameClass,
		chroma.NameConstant,
		chroma.NameDecorator,
		chroma.NameEntity,
		chroma.NameException,
		chroma.NameFunction,
		chroma.NameFunctionMagic,
		chroma.NameKeyword,
		chroma.NameOperator:
		return "text-blue-600 dark:text-blue-400"

	case chroma.NameOther:
		return "text-muted-foreground"

	case chroma.Name,
		chroma.NameBuiltin,
		chroma.NameBuiltinPseudo,
		chroma.NameLabel,
		chroma.NameNamespace,
		chroma.NamePseudo,
		chroma.NameProperty,
		chroma.NameTag,
		chroma.NameVariable,
		chroma.NameVariableAnonymous,
		chroma.NameVariableClass,
		chroma.NameVariableGlobal,
		chroma.NameVariableInstance,
		chroma.NameVariableMagic:
		return "text-neutral-900 dark:text-neutral-100"

	case chroma.LiteralString,
		chroma.LiteralStringAffix,
		chroma.LiteralStringAtom,
		chroma.LiteralStringBacktick,
		chroma.LiteralStringBoolean,
		chroma.LiteralStringChar,
		chroma.LiteralStringDelimiter,
		chroma.LiteralStringDoc,
		chroma.LiteralStringDouble,
		chroma.LiteralStringEscape,
		chroma.LiteralStringHeredoc,
		chroma.LiteralStringInterpol,
		chroma.LiteralStringName,
		chroma.LiteralStringOther,
		chroma.LiteralStringRegex,
		chroma.LiteralStringSingle,
		chroma.LiteralStringSymbol:
		return "text-emerald-700 dark:text-emerald-300"

	case chroma.Literal,
		chroma.LiteralNumber,
		chroma.LiteralDate,
		chroma.LiteralOther,
		chroma.LiteralNumberBin,
		chroma.LiteralNumberFloat,
		chroma.LiteralNumberHex,
		chroma.LiteralNumberInteger,
		chroma.LiteralNumberIntegerLong,
		chroma.LiteralNumberOct,
		chroma.LiteralNumberByte:
		return "text-orange-600 dark:text-orange-400"

	case chroma.Operator, chroma.OperatorWord:
		return "text-neutral-400 dark:text-neutral-600"

	case chroma.Punctuation:
		return "text-muted-foreground"

	case chroma.Comment,
		chroma.CommentHashbang,
		chroma.CommentMultiline,
		chroma.CommentSingle,
		chroma.CommentSpecial,
		chroma.CommentPreproc,
		chroma.CommentPreprocFile:
		return "text-neutral-400 dark:text-neutral-600"

	case chroma.GenericEmph:
		return "italic"
	case chroma.GenericStrong:
		return "font-bold"
	case chroma.GenericUnderline:
		return "underline"
	default:
		return ""
	}
}

type ctxwriter struct {
	io.Writer
	ctx context.Context
}

var htmlformatter = chroma.FormatterFunc(func(w io.Writer, style *chroma.Style, iterator chroma.Iterator) error {
	for token := range iterator.Stdlib() {
		class := gettokclass(token.Type)
		el := html.Span(
			html.If(len(class) > 0, html.Class(class)),

			html.Text(token.Value),
		)
		ctx := context.Background()
		ctxwer, ok := w.(*ctxwriter)
		if ok {
			ctx = ctxwer.ctx
		}

		for chunk := range el.Chunks() {
			if err := html.Render(chunk, ctx, w); err != nil {
				return err
			}
		}
	}

	return nil
})

func CodeComponent(language, code, name string, accsessories bool) html.Node {
	lexer := lexers.Get(language)
	if lexer == nil {
		lexer = lexers.Fallback
	}
	lexer = chroma.Coalesce(lexer)
	iterator, err := lexer.Tokenise(nil, code)
	if err != nil {
		log.Fatal(err)
	}

	var buf = new(bytes.Buffer)
	if err := htmlformatter.Format(&ctxwriter{Writer: buf, ctx: context.Background()}, styles.Fallback, iterator); err != nil {
		log.Fatal(err)
	}

	return html.Div(
		html.Attr("x-data", "{}"),
		html.Class("text-sm p-4 bg-background border rounded-md"),

		html.If(accsessories,
			html.Div(
				html.Class("flex justify-between gap-4 border-b pb-1.5 mb-1.5 text-muted-foreground"),

				html.P(html.Class("text-xs"), html.Text(name)),
				html.Button(
					x.On("click", "navigator.clipboard.writeText($refs.code.innerText)"),

					lucide.Copy(html.Class("size-4")),
				),
			),
		),
		html.Code(
			x.Ref("code"),

			html.Pre(html.Class("whitespace-pre-wrap"), html.RawUnsafe(buf.String())),
		),
	)
}
