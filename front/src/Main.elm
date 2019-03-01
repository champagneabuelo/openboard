module Main exposing (Model, Msg(..), init, linkBtn, main, theme, update, view)

import Browser
import Css exposing (..)
import Html
import Html.Styled exposing (..)
import Html.Styled.Attributes exposing (css, href, src)
import Html.Styled.Events exposing (onClick)



---- MODEL ----


type alias Model =
    {}


init : ( Model, Cmd Msg )
init =
    ( {}, Cmd.none )



---- UPDATE ----


type Msg
    = NoOp


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    ( model, Cmd.none )



---- VIEW ----


view : Model -> Html Msg
view model =
    card []
        [ heading [] [ text "openboard!" ]
        , linkBtn [ href "https://github.com/rtfeldman/elm-css" ] [ text "Learn more about elm css" ]
        , linkBtn [ href "hhttps://github.com/champagneabuelo/openboard" ] [ text "githhub" ]
        ]



---- PROGRAM ----


main : Program () Model Msg
main =
    Browser.element
        { view = view >> toUnstyled
        , init = \_ -> init
        , update = update
        , subscriptions = always Sub.none
        }


gridUnit =
    8


card =
    styled div
        [ borderRadius (px 4)
        , padding (px (gridUnit * 4))
        , boxShadow4 (px 1) (px 1) (px 5) theme.dark
        , margin (px (8 * gridUnit))
        ]


heading =
    styled h1
        [ h1Style ]


h1Style : Style
h1Style =
    Css.batch
        [ fontFamilies [ "Palatino Linotype", "Georgia", "serif" ]
        , fontSize (6 * gridUnit |> px)
        , fontWeight bold
        , marginTop (px 0)
        ]


{-| A reusable button which has some styles pre-applied to it.
-}
linkBtn : List (Attribute msg) -> List (Html msg) -> Html msg
linkBtn =
    styled a
        [ padding4 (px 8) (px 16) (px 8) (px 16)
        , color (rgb 250 250 250)
        , hover
            [ backgroundColor theme.hover
            , textDecoration underline
            ]
        , display block
        , margin (px 10)
        , backgroundColor theme.primary
        , cursor pointer
        , border (px 0)
        , borderRadius (px 3)
        , fontSize (Css.em 1)
        ]


{-| A plain old record holding a couple of theme colors.
-}
theme =
    { primary = hex "55af6a"
    , secondary = rgb 250 240 230
    , hover = hex "3ebc5b"
    , dark = hex "999"
    }
