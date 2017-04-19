import 'dart:async';

import 'package:angular2/core.dart';

import 'package:landgrab/component/board/component.dart';
import 'package:landgrab/component/players_choice_form/component.dart';
import 'package:landgrab/component/rules/component.dart';
import 'package:landgrab/model/state.dart';
import 'package:landgrab/model/rules.dart';
import 'package:landgrab/model/player.dart';
import 'package:landgrab/service/rules.dart';
import 'package:landgrab/service/players.dart';
import 'package:landgrab/service/state.dart';

@Component(
  selector: 'game',
  templateUrl: 'template.html',
  directives: const [
    RulesComponent,
    PlayersChoiceFormComponent,
    BoardComponent
  ],
  providers: const [StateService],
)
class GameComponent implements OnInit {
  String status;

  RulesService _rulesService;

  PlayersService _playersService;

  StateService _stateService;

  Rules rules;

  final List<Player> players = [];

  Player player1;

  Player player2;

  State state;

  GameComponent(this._rulesService, this._playersService, this._stateService);

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
    } catch (e) {
      status = 'Too many requests';
      print(e);
    }
  }

  reset() {
    state = null;
    player1 = players.first;
    player2 = players.first;
  }

  Future next() async {
    if (state == null) return;
    try {
      status = '';
      state.startTimer();
      State newState = await _stateService.next(state);
      state.stopTimer();
      state = newState;
    } catch (error) {
      state.stopTimer();
      status = 'Too many requests';
      print(error);
    }
  }

  isWinner(PlayerID id) => id != PlayerID.noPlayer;
}
