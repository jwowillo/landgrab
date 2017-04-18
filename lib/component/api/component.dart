import 'dart:async';

import 'package:angular2/core.dart';

import 'package:landgrab/component/resource/component.dart';
import 'package:landgrab/component/controller_description/component.dart';
import 'package:landgrab/model/api.dart';
import 'package:landgrab/pipe/pretty_json.dart';
import 'package:landgrab/service/schema.dart';

@Component(
  selector: 'api',
  templateUrl: 'template.html',
  directives: const [ResourceComponent, ControllerDescriptionComponent],
  providers: const [SchemaService],
  pipes: const [PrettyJSONPipe],
)
class APIComponent implements OnInit {
  final List<Resource> resources = [];
  final List<Action> actions = [];

  SchemaService _service;

  APIComponent(this._service);

  @override
  Future ngOnInit() async {
    try {
      resources.addAll(await _service.resources());
      actions.addAll(await _service.actions());
    } catch (error) {
      print(error);
    }
  }
}
