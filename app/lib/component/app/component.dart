import 'package:angular2/core.dart';

import 'package:landgrab/component/game/component.dart';

@Component(
  selector: 'app',
  templateUrl: 'template.html',
  directives: const [GameComponent],
)
class AppComponent {}
