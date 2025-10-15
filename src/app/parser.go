package app

import (
	"embed"
	"log"
	"os"
	"strings"

	"github.com/canpacis/pacis-docs/src/components/code"
	. "github.com/canpacis/pacis/html"
	"github.com/canpacis/pacis/server"
	parser "github.com/sivukhin/godjot/djot_parser"
)

func BuildMarkup(node parser.TreeNode[parser.DjotNode]) Node {
	var element *Element

	switch node.Type {
	case parser.ParagraphNode:
		element = P(Class("inline leading-relaxed"))
	case parser.LinkNode:
		element = A(
			Class("text-sky-400 underline inline break-all"),
			Href(node.Attributes.Get("href")),
		)
		if strings.HasPrefix(node.Attributes.Get("href"), "https://") {
			element.SetAttribute("target", "_blank")
		}
	case parser.UnorderedListNode:
		element = Ul(Class("list-disc list-inside"))
	case parser.OrderedListNode:
		element = Ol(Class("list-decimal list-inside"))
	case parser.ListItemNode:
		element = Li(Class("my-2"))
	case parser.HeadingNode:
		switch node.Attributes.Get(parser.HeadingLevelKey) {
		case "#":
			element = H1(Class("text-3xl font-semibold mb-2"))
		case "##":
			element = H2(Class("text-2xl font-semibold mb-2"))
		case "###":
			element = H3(Class("text-xl font-semibold mb-2"))
		default:
			element = H4(Class("text-lg font-semibold mb-2"))
		}
	case parser.TableNode:
		element = Table(Class("w-full caption-bottom text-sm"))
	case parser.TableHeaderNode:
		element = Th(Class("h-10 px-2 text-left align-middle font-medium text-muted-foreground [&:has([role=checkbox])]:pr-0 [&>[role=checkbox]]:translate-y-[2px]"))
	case parser.TableRowNode:
		element = Tr(Class("border-b transition-colors hover:bg-muted/50 data-[state=selected]:bg-muted"))
	case parser.TableCellNode:
		element = Td(Class("p-2 align-middle [&:has([role=checkbox])]:pr-0 [&>[role=checkbox]]:translate-y-[2px]"))
	case parser.SectionNode:
		element = Section(Class("flex flex-col gap-4"), ID(node.Attributes.Get("id")))
	case parser.StrongNode:
		element = Span(Class("font-semibold"))
	case parser.SuperscriptNode:
		element = Sup()
	case parser.SubscriptNode:
		element = Sub()
	case parser.FootnoteDefNode:
		element = Span()
	case parser.EmphasisNode:
		element = Span(Class("italic"))
	case parser.ImageNode:
		element = Img(Class("w-full rounded-lg"), Src(node.Attributes.Get("src")))
	case parser.QuoteNode:
		element = Blockquote(Class("border-l-5 border-accent pl-2 py-2"))
	case parser.VerbatimNode:
		element = Span(Class("bg-accent rounded-sm text-sm py-1 px-2 font-mono"))
	case parser.ThematicBreakNode:
		element = Hr()
	case parser.CodeNode:
		lang := node.Attributes.Get("$CodeLangKey")
		noaccessory := node.Attributes.Get("noaccessory")
		return code.New(
			lang,
			string(node.FullText()),
			node.Attributes.Get("file"),
			lang == "go" && noaccessory != "true",
		)
	case parser.TextNode:
		r := strings.NewReplacer("&ldquo;", "\"", "&rdquo;", "\"", "&lsquo;", "'", "&rsquo;", "'", "&hellip;", "...")
		return Text(r.Replace(string(node.Text)))
	default:
		nodes := Fragment()
		for _, child := range node.Children {
			nodes = append(nodes, BuildMarkup(child))
		}
		return nodes
	}

	class := node.Attributes.Get("class")
	if len(class) > 0 {
		element.AddClass(class)
	}
	for _, child := range node.Children {
		element.AppendNode(BuildMarkup(child))
	}
	return element
}

func ExtractMetadata(node parser.TreeNode[parser.DjotNode]) (string, string) {
	var title string
	var desc string
	node.Traverse(func(node parser.TreeNode[parser.DjotNode]) {
		switch node.Type {
		case parser.HeadingNode:
			if (node.Attributes.Get(parser.HeadingLevelKey) == "#") && len(title) == 0 {
				title = string(node.FullText())
			}
		case parser.ParagraphNode:
			if len(desc) == 0 {
				desc = string(node.FullText())
				if len(desc) > 260 {
					desc = desc[:260] + "..."
				}
			}
		}
	})

	return title, desc
}

var docsfs embed.FS

func SetDocsFS(fs embed.FS) {
	docsfs = fs
}

func GetDocFile(env server.Environment, name string) []byte {
	if env == server.Dev {
		file, err := os.ReadFile("src/app/docs/" + name + ".md")
		if err != nil {
			log.Fatal(err)
		}
		return file
	}
	file, err := docsfs.ReadFile("src/app/docs/" + name + ".md")
	if err != nil {
		log.Fatal(err)
	}
	return file
}
