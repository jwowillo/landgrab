import 'dart:async';

import 'package:angular2/core.dart';

import 'package:landgrab/landgrab.dart';

@Injectable()
class SchemaService {

  final APIService _service;

  static List<Resource> _resources = [];

  static List<Action> _actions = [];

  SchemaService(this._service);

  Future<List<Resource>> resources() async {
    if (!_resources.isEmpty) {
      return _resources;
    }
    Map<String, dynamic> json = await _service.request('/schema');
    for (Map<String, dynamic> resource in json['data']['resources']) {
      _resources.add(mapToResource(resource));
    }
    return _resources;
  }

  Future<List<Action>> actions() async {
    if (!_actions.isEmpty) {
      return _actions;
    }
    Map<String, dynamic> json = await _service.request('/schema');
    for (Map<String, dynamic> action in json['data']['actions']) {
      _actions.add(mapToAction(action));
    }
    return _actions;
  }
}
