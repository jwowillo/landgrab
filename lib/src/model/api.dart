class Resource {
  final String name;
  final String description;
  final Object structure;

  const Resource(this.name, this.description, this.structure);
}

class Action {
  final String path;
  final ControllerDescription description;

  const Action(this.path, this.description);
}

class ControllerDescription {
  final MethodDescription getMethod;
  final MethodDescription postMethod;
  final MethodDescription putMethod;
  final MethodDescription deleteMethod;
  final MethodDescription optionsMethod;
  final MethodDescription headMethod;
  final MethodDescription connectMethod;

  const ControllerDescription(
      {this.getMethod,
      this.postMethod,
      this.putMethod,
      this.deleteMethod,
      this.optionsMethod,
      this.headMethod,
      this.connectMethod});
}

class MethodDescription {
  final String name;
  final Map<String, String> urlArguments;
  final Map<String, String> formArguments;
  final String response;
  final String authentication;
  final String limiting;

  const MethodDescription(this.name,
      {this.urlArguments,
      this.formArguments,
      this.response,
      this.authentication,
      this.limiting});
}
