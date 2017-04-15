import 'package:angular2/core.dart';
import 'package:angular2/router.dart';

import 'package:landgrab/component/api/component.dart';
import 'package:landgrab/component/game/component.dart';

@Component(
  selector: 'app',
  templateUrl: 'template.html',
  directives: const [ROUTER_DIRECTIVES, GameComponent, APIComponent],
  providers: const [ROUTER_PROVIDERS],
)
@RouteConfig(const [
  const Route(path: '/', name: 'Game', component: GameComponent),
  const Route(path: '/api', name: 'API', component: APIComponent),
])
class AppComponent {}
