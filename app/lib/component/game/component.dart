import 'dart:async';

import 'package:angular2/core.dart';
import 'package:angular2/router.dart';

import 'package:landgrab/model/board.dart';
import 'package:landgrab/model/player.dart';
import 'package:landgrab/model/state.dart';
import 'package:landgrab/pipe/player_id_to_string.dart';
import 'package:landgrab/pipe/to_lower_no_space.dart';
import 'package:landgrab/service/state.dart';

/// GameComponent contains a landgrab game with Players already chosen.
/// /// Another parameter, named wait, can be passed. If true, the game will prompt
/// the user before proceeding to the next turn.
///
/// At the end of the game, a link will be provided to go back to the
/// GameFormComponent.
@Component(
  selector: 'game',
  templateUrl: 'template.html',
  styleUrls: const ['styles.css'],
  providers: const [StateService],
  directives: const [ROUTER_DIRECTIVES],
  pipes: const [PlayerIDToStringPipe, ToLowerNoSpacePipe],
)
class GameComponent implements OnInit {
  /// _service for getting the initial and next States.
  final StateService _service;

  // _route which triggered the GameComponent.
  final RouteParams _route;

  /// _state of the game.
  State _state;

  /// GameComponent constructor initializes the StateService and RouteParams.
  GameComponent(this._service, this._route);

  /// ngOnInit loads the initial State.
  @override
  Future ngOnInit() async {
    Player player1 = new Player(_route.get('player1'));
    Player player2 = new Player(_route.get('player2'));
    try {
      _state = await _service.initial(player1, player2);
    } catch (error) {
      print(error);
    }
  }

  /// next loads the next State.
  Future next() async {
    try {
      _state = await _service.next(_state);
    } catch (error) {
      print(error);
    }
  }

  /// state getter.
  State get state => _state;

  /// isPiece returns true iff the Piece isn't a NO_PIECE.
  isPiece(Piece piece) => piece != NO_PIECE;

  /// isPlayer1 returns true iff the Piece is owned by PlayerID.player1.
  isPlayer1(Piece p) => _state.playerForPiece(p) == PlayerID.player1;

  /// isPlayer1 returns true iff the Piece is owned by PlayerID.player2.
  isPlayer2(Piece p) => _state.playerForPiece(p) == PlayerID.player2;

  /// isWinner returns true iff the PlayerID isn't PlayerID.noPlayer.
  isWinner(PlayerID id) => id != PlayerID.noPlayer;
}
