package app

import . "github.com/canpacis/pacis/html"

var previews = map[string][]Node{}

func Preview(child Node) Node {
	return Div(
		Class("data-[tab=code]:border-code relative rounded-lg border data-[chrome-less-on-mobile=true]:border-0 sm:data-[chrome-less-on-mobile=true]:border md:-mx-1"),

		Div(
			Class("preview flex w-full justify-center data-[align=center]:items-center data-[align=end]:items-end data-[align=start]:items-start h-[450px] p-10"),

			child,
		),
	)
}
