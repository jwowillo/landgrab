import 'dart:async';

import 'package:angular2/core.dart';

import 'package:landgrab/service/api.dart';
import 'package:landgrab/service/convert.dart';
import 'package:landgrab/model/api.dart';

@Injectable()
class SchemaService {
  Future<List<Resource>> resources() async {
    Map<String, dynamic> json = await api('/schema');
    List<Resource> resources = [];
    for (Map<String, dynamic> resource in json['data']['resources']) {
      resources.add(mapToResource(resource));
    }
    return resources;
  }

  Future<List<Action>> actions() async {
    Map<String, dynamic> json = await api('/schema');
    List<Action> actions = [];
    for (Map<String, dynamic> action in json['data']['actions']) {
      actions.add(mapToAction(action));
    }
    return actions;
  }
}
