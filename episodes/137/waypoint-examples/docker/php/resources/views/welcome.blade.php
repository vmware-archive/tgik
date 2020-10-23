<!DOCTYPE html>
<html lang="{{ str_replace('_', '-', app()->getLocale()) }}">
    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">

        <title>Waypoint Laravel Example</title>
        <link rel="stylesheet" type="text/css" href="{{ asset('css/main.css') }}" />
    </head>
    <body>
        <div class="container">
            <header>
                <a href="https://waypointproject.io" class="logo">
                <img src="{{ asset('images/logo.svg') }}" alt="Logo" />
                </a>
            </header>
            <section class="content">
                <div class="language-icon">
                <img src="{{ asset('images/language.svg') }}" alt="PHP Icon" width="65px" />
                </div>
                <h1>This Laravel app was deployed with Waypoint.</h1>
                <p>
                    Try making a change to this text locally and run <code>waypoint up</code> again to see it.
                </p>
                <p>
                Read the <a href="https://waypointproject.io/docs">documentation</a> for more about Waypoint.
                </p>
            </section>
            <footer>
                <a href="https://hashicorp.com" class="hashi">
                <img src="{{ asset('images/hashi.svg') }}" alt="HashiCorp" />
                </a>
            </footer>
        </div>
    </body>
</html>
