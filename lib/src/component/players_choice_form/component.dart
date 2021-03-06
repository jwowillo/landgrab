import 'package:angular2/core.dart';

import 'package:landgrab/landgrab.dart';

/// FormComponent allows options to be selected before routing to the
/// GameComponent.
///
/// These options include what the Player 1 and Player 2 are and if the
/// GameComponent should prompt the user before continuing to the next turn.
@Component(
  selector: 'players-choice-form',
  templateUrl: 'template.html',
  directives: const [PlayerChoiceFormComponent],
)
class PlayersChoiceFormComponent {
  @Input()
  Set<Player> players;

  @Output()
  final EventEmitter<Player> player1Chosen = new EventEmitter();

  @Output()
  final EventEmitter<Player> player2Chosen = new EventEmitter();

  /// player1's PlayerID.
  final PlayerID player1ID = PlayerID.player1;

  /// player2's PlayerID.
  final PlayerID player2ID = PlayerID.player2;

  /// player1 chosen.
  Player player1;

  /// player2 chosen.
  Player player2;
}
