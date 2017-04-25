import 'dart:async';

import 'package:angular2/core.dart';

import 'package:landgrab/landgrab.dart';

/// StateService is used to get initial States for games with provided Players
/// and get the next States for given States.
@Injectable()
class MovesService {
  final APIService _service;

  MovesService(this._service);

  /// next State for the given State.
  Future<List<Move>> moves(State s) async {
    String serialized = encode(stateToMap(s));
    Map<String, dynamic> json =
        await _service.request('/moves', query: {'state': serialized});
    return mapToMoves(json['data']);
  }
}
