import 'package:angular2/core.dart';
import 'package:angular2/router.dart';

import 'package:landgrab/component/game/component.dart';
import 'package:landgrab/component/game_form/component.dart';

/// AppComponent establishes routing between the initial GameFormComponent and
/// the next GameComponent.
///
/// AppComponent also shows diagnostic messages about the API server.
@Component(
  selector: 'app',
  templateUrl: 'template.html',
  directives: const [ROUTER_DIRECTIVES],
  providers: const [ROUTER_PROVIDERS],
)
@RouteConfig(const [
  const Route(path: '/', name: 'GameForm', component: GameFormComponent),
  const Route(path: '/game', name: 'Game', component: GameComponent),
])
class AppComponent {}
