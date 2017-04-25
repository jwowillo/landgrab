import 'dart:async';

import 'package:angular2/core.dart';

import 'package:landgrab/landgrab.dart';

@Component(
  selector: 'game',
  templateUrl: 'template.html',
  directives: const [
    RulesComponent,
    PlayersChoiceFormComponent,
    BoardComponent
  ],
  providers: const [StateService, PlayersService, RulesService, MovesService],
)
class GameComponent implements OnInit {
  String status;

  RulesService _rulesService;

  PlayersService _playersService;

  StateService _stateService;

  MovesService _movesService;

  Rules rules;

  final List<Player> players = [];

  Player player1;

  Player player2;

  State state;

  List<Move> moves = [];

  GameComponent(this._rulesService, this._playersService, this._stateService,
      this._movesService);

  @override
  Future ngOnInit() async {
    try {
      status = '';
      rules = await _rulesService.rules();
      players.addAll(await _playersService.players());
      players.sort((Player a, Player b) => a.name.compareTo(b.name));
      player1 = players.first;
      player2 = players.first;
    } catch (error) {
      status = 'Too many requests';
      print(error);
    }
  }

  Future start() async {
    if (player1 == null || player2 == null) return;
    try {
      status = '';
      state = await _stateService.initial(player1, player2);
      moves = await _movesService.moves(state);
    } catch (e) {
      status = 'Too many requests';
      print(e);
    }
  }

  reset() {
    state = null;
    player1 = players.first;
    player2 = players.first;
    moves = [];
  }

  Future next() async {
    if (state == null) return;
    try {
      status = '';
      state.startTimer();
      State newState = await _stateService.next(state);
      state.stopTimer();
      state = newState;
      player1.arguments = null;
      player2.arguments = null;
      state.player1.arguments = null;
      state.player2.arguments = null;
      moves = await _movesService.moves(state);
    } catch (error) {
      state.stopTimer();
      status = 'Too many requests';
      print(error);
    }
  }

  isWinner(PlayerID id) => id != PlayerID.noPlayer;

  changed(Map<String, dynamic> arguments) {
    if (state == null) {
      return;
    }
    if (state.currentPlayer == PlayerID.player1) {
      player1.arguments = arguments;
      state.player1.arguments = arguments;
    }
    if (state.currentPlayer == PlayerID.player2) {
      player2.arguments = arguments;
      state.player2.arguments = arguments;
    }
  }
}
