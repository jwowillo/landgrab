import 'dart:async';

import 'package:angular2/core.dart';

import 'package:landgrab/component/board/component.dart';
import 'package:landgrab/component/players_choice_form/component.dart';
import 'package:landgrab/component/rules/component.dart';
import 'package:landgrab/model/state.dart';
import 'package:landgrab/model/rules.dart';
import 'package:landgrab/model/player.dart';
import 'package:landgrab/service/players.dart';
import 'package:landgrab/service/rules.dart';
import 'package:landgrab/service/state.dart';

@Component(
  selector: 'game',
  templateUrl: 'template.html',
  directives: const [
    RulesComponent,
    PlayersChoiceFormComponent,
    BoardComponent
  ],
  providers: const [RulesService, PlayersService, StateService],
)
class GameComponent implements OnInit {
  RulesService _rulesService;

  PlayersService _playersService;

  StateService _stateService;

  Rules rules;

  final Set<Player> players = new Set();

  Player player1;

  Player player2;

  State state;

  GameComponent(this._rulesService, this._playersService, this._stateService);

  @override
  Future ngOnInit() async {
    try {
      rules = await _rulesService.rules();
      players.addAll(await _playersService.players());
      player1 = players.first;
      player2 = players.first;
    } catch (error) {
      print(error);
    }
  }

  Future start() async {
    if (player1 == null || player2 == null) return;
    state = await _stateService.initial(player1, player2);
  }

  reset() {
    state = null;
    player1 = players.first;
    player2 = players.first;
  }

  Future next() async {
    if (state == null) return;
    try {
      state.startTimer();
      State newState = await _stateService.next(state);
      state.stopTimer();
      state = newState;
    } catch (error) {
      print(error);
    }
  }

  isWinner(PlayerID id) => id != PlayerID.noPlayer;
}
