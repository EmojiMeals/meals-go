# meals-go ![emoji? yes](https://img.shields.io/badge/emoji%3F-%F0%9F%91%8D-brightgreen) ![Tests](https://github.com/EmojiMeals/meals-go/workflows/Tests/badge.svg?branch=master)

Find what you can cook with your emojis

## get started

install the package:

```console
go get github.com/EmojiMeals/meals-go
```

then start using it:

```console
import "github.com/EmojiMeals/meals-go"

cookbook := meals.NewCookbook("path-to-recipes.json")

cookbook.Mealify("🍞","🍅","🧀")
```

## Usage 

You can start checking for recipes using `cookbook.Mealify()`.

```golang
meals.Mealify("🌚","🍰") #=> "🥮"

# it is not order dependent 💯💯IQ 🧠💥
meals.Mealify("🍰", "🌚") #=> "🥮"

# Use more than two items!
meals.Mealify("🍞","🍅","🧀") #=> "🍕"

# After all this eating, I need a drink
meals.Mealify("🍶","🍾","🍷","🍸","🍶","🍹","🍺","🍻","🥂","🍾","🥃") #=> "🤮"
```

## Is it thread safe

Yes!

## Can I donate to the project?
[Yes, here](https://www.buymeacoffee.com/emoji)

## License

This project is licensed under the [MIT License](https://github.com/EmojiMeals/meals-go/blob/master/LICENSE).
