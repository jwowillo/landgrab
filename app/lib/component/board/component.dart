import 'package:angular2/core.dart';

import 'package:landgrab/model/board.dart';
import 'package:landgrab/model/player.dart';
import 'package:landgrab/model/state.dart';
import 'package:landgrab/model/rules.dart';
import 'package:landgrab/pipe/player_id_to_string.dart';
import 'package:landgrab/pipe/to_lower_no_space.dart';

/// BoardComponent contains a landgrab board with Players already chosen.
/// /// Another parameter, named wait, can be passed. If true, the board will prompt
/// the user before proceeding to the next turn.
///
/// At the end of the board, a link will be provided to go back to the
/// GameFormComponent.
@Component(
  selector: 'board',
  templateUrl: 'template.html',
  styleUrls: const ['styles.css'],
  pipes: const [PlayerIDToStringPipe, ToLowerNoSpacePipe],
)
class BoardComponent {
  @Input()
  Rules rules;

  /// state of the board.
  @Input()
  State state;

  List<List<Piece>> get grid {
    List<List<Piece>> g = [];
    if (rules == null) return g;
    for (int i = 0; i < rules.boardSize; i++) {
      List<Piece> row = [];
      for (int j = 0; j < rules.boardSize; j++) {
        if (state == null) {
          row.add(NO_PIECE);
        } else {
          row.add(state.pieceForCell(new Cell(i, j)));
        }
      }
      g.add(row);
    }
    return g;
  }

  /// isPiece returns true iff the Piece isn't a NO_PIECE.
  isPiece(Piece piece) => piece != NO_PIECE;

  /// isPlayer1 returns true iff the Piece is owned by PlayerID.player1.
  isPlayer1(Piece p) => state.playerForPiece(p) == PlayerID.player1;

  /// isPlayer1 returns true iff the Piece is owned by PlayerID.player2.
  isPlayer2(Piece p) => state.playerForPiece(p) == PlayerID.player2;

  /// isWinner returns true iff the PlayerID isn't PlayerID.noPlayer.
  isWinner(PlayerID id) => id != PlayerID.noPlayer;

  String get timeRemaining {
    int ms = state.timeRemaining.inMilliseconds;
    int s = (ms / Duration.MILLISECONDS_PER_SECOND.toDouble()).toInt();
    ms -= s * Duration.MILLISECONDS_PER_SECOND;
    ms = (ms / 100).toInt();
    return '$s.$ms seconds';
  }
}
