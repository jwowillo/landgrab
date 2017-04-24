import 'package:angular2/angular2.dart';

import 'package:landgrab/landgrab.dart';

/// PlayerIDToStringPipe is a pipe which transforms a PlayerID to a String
/// representation of that PlayerID.
@Pipe(name: 'playerIDToString')
class PlayerIDToStringPipe extends PipeTransform {
  /// transform the PlayerID into a String representation of the PlayerID.
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
