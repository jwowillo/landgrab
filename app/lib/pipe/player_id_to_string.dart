import 'package:angular2/angular2.dart';

import 'package:landgrab/model/player.dart';

@Pipe(name: 'playerIDToString')
class PlayerIDToStringPipe extends PipeTransform {
  String transform(PlayerID id) {
    if (id == PlayerID.player1) {
      return 'Player 1';
    }
    if (id == PlayerID.player2) {
      return 'Player 2';
    }
    return '';
  }
}
