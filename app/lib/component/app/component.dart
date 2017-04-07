import 'package:angular2/core.dart';

import 'package:landgrab/component/player_choice_form/component.dart';
import 'package:landgrab/model/player.dart';

// AppComponent shows the GameFormComponent until it is submitted and then shows
// the GameComponent until the game is complete.
@Component(
  selector: 'app',
  templateUrl: 'template.html',
  directives: const [PlayerChoiceFormComponent],
)
class AppComponent {
  // player1's PlayerID.
  final PlayerID player1 = PlayerID.player1;
  // player2's PlayerID.
  final PlayerID player2 = PlayerID.player2;
}
