package ui

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// BUTTONS
func ButtonUI(children ...Node) Node {
	return Button(
		InlineStyle(`
			$me {
				cursor: pointer;
				background-color: $color(red-700);
				color: $color(neutral-50);
				padding: $1.5 $8 $1.5 $8;
				font-size: var(--text-sm);
				border-radius: var(--radius-xs);
				font-weight: var(--font-weight-semibold);
			}

			$me:hover {
				background: $color(red-800);
			}
		`),
		Group(children),
	)
}

func ButtonUIOutline(children ...Node) Node {
	return ButtonUI(
		InlineStyle(`
			$me {
				background: none;
				box-sizing: border-box;
                box-shadow: 0 0 0 0, inset 0 0 0 1px $color(gray-400);
				color: $color(gray-400);
			}

			$me:hover {
				color: $color(white);
				box-shadow: none;
				background: $color(gray-400);
			}
		`),
		Group(children),
	)
}
func ButtonUISuccess(children ...Node) Node {
	return ButtonUI(
		InlineStyle(`
			$me {
				background: $color(green-700);
				color: $color(white);
			}

			$me:hover {
				background: $color(green-800);
			}
		`),
		Group(children),
	)
}