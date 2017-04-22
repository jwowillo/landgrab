import 'dart:async';

import 'package:angular2/core.dart';

import 'package:landgrab/model/rules.dart';
import 'package:landgrab/service/api.dart';
import 'package:landgrab/service/convert.dart';

@Injectable()
class RulesService {
  final APIService _service;

  static Rules _rules;

  RulesService(this._service);

  Future<Rules> rules() async {
    if (_rules != null) {
      return _rules;
    }
    Map<String, dynamic> json = await _service.request('/rules');
    _rules = mapToRules(json['data']['rules']);
    return _rules;
  }
}
