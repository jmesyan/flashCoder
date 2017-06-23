package main

import (
	"fmt"
	"goreact/vendor"
)

func main() {
	v := vendor.NewRenderer([]string{"assets/demo1.js"}).
		RunCmd(`
			var component = React.createElement(CommentBox, { foo: 'bar' });
			React.renderToString(component);
		`)
	fmt.Printf("\n%v\n", v)

	v = vendor.NewRenderer([]string{"assets/demo2.js"}).
		RunCmd(`
			var component = React.createElement(HelloWorld, { foo: 'bar' });
			React.renderToString(component);
		`)
	fmt.Printf("\n%v\n", v)

	v = vendor.NewRenderer([]string{"assets/demo3.js"}).
		RunCmd(`
			var data = [
				{"id": 0, "author": "Anonymous", "text": "This is a comment"},
				{"id": 1, "author": "Anonymous", "text": "This is another comment"},
			]
			var component = React.createElement(CommentBox, {data : data});
			React.renderToString(component);
		`)
	fmt.Printf("\n%v\n", v)
}
