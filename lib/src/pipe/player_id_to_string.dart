import 'package:angular2/angular2.dart';

import 'package:landgrab/landgrab.dart';

/// PlayerIDToStringPipe is a pipe which transforms a PlayerID to a String
/// representation of that PlayerID.
@Pipe(name: 'playerIDToString')
class PlayerIDToStringPipe extends PipeTransform {
  /// transform the PlayerID into a String representation of the PlayerID.
  String transform(PlayerID id) {
    String x = playerIDToString(id);
    return x[0].toUpperCase() + x.substring(1);
  }
}
