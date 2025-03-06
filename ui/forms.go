package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// FORMS
func FormInput(children ...Node) Node {
	return Input(
		InlineStyle(`
			$me {
				background-color: $color(white);
				padding: $2;
				display: block;
				width: 100%;
				border: 0;
				color: $color(neutral-900);
				box-shadow: var(--shadow-xs);
			}

			@media $sm {
				$me {
					font-size: var(--text-sm);
				}
			}
		`),
		Group(children),
	)
}

func FormSelect(children ...Node) Node {
	return Select(
		InlineStyle(`
			$me {
				padding: $3;
				background-color: $color(white);
				display: block;
				width: 100%;
				border: 0;
				color: $color(neutral-900);
				box-shadow: var(--shadow-xs);
			}
			@media $sm {
				$me {
					font-size: var(--text-sm);
				}
			}
		`),
		Group(children),
	)
}

func FormTextarea(children ...Node) Node {
	return Textarea(
		InlineStyle(`
			$me {
				display: block;
				padding: $3;
				width: 100%;
				font-size: var(--text-sm);
				color: $color(neutral-900);
				background-color: $color(white);
				box-shadow: var(--shadow-xs);
			}
		`),
		Group(children),
	)
}

func FormLabel(children ...Node) Node {
	return Label(Class("text-neutral-900 text-sm"), Group(children))
}