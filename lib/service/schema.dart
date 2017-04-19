import 'dart:async';

import 'package:angular2/core.dart';

import 'package:landgrab/service/api.dart';
import 'package:landgrab/service/convert.dart';
import 'package:landgrab/model/api.dart';

@Injectable()
class SchemaService {
  List<Resource> _resources = [];

  List<Action> _actions = [];

  Future<List<Resource>> resources() async {
    if (!_resources.isEmpty) {
      return _resources;
    }
    Map<String, dynamic> json = await api('/schema');
    for (Map<String, dynamic> resource in json['data']['resources']) {
      _resources.add(mapToResource(resource));
    }
    return _resources;
  }

  Future<List<Action>> actions() async {
    if (!_actions.isEmpty) {
      return _actions;
    }
    Map<String, dynamic> json = await api('/schema');
    for (Map<String, dynamic> action in json['data']['actions']) {
      _actions.add(mapToAction(action));
    }
    return _actions;
  }
}
