import 'package:angular2/core.dart';
import 'package:angular2/router.dart';


/// GameComponent contains a landgrab game with Players already chosen.
///
/// Another parameter, named wait, can be passed. If true, the game will prompt
/// the user before proceeding to the next turn.
///
/// At the end of the game, a link will be provided to go back to the
/// GameFormComponent.
@Component(
  selector: 'game',
  template: '''
  <h2>GAME</h2>
  <a [routerLink]="['GameForm']">Back</a>
  ''',
  directives: const [ROUTER_DIRECTIVES],
)
class GameComponent {}
