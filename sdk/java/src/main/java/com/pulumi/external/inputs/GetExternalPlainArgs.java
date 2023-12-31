// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.external.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetExternalPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetExternalPlainArgs Empty = new GetExternalPlainArgs();

    @Import(name="programs", required=true)
    private List<String> programs;

    public List<String> programs() {
        return this.programs;
    }

    /**
     * A map of string values to pass to the external program as the query arguments. If not supplied, the program will receive an empty object as its input.
     * 
     */
    @Import(name="query")
    private @Nullable Map<String,String> query;

    /**
     * @return A map of string values to pass to the external program as the query arguments. If not supplied, the program will receive an empty object as its input.
     * 
     */
    public Optional<Map<String,String>> query() {
        return Optional.ofNullable(this.query);
    }

    /**
     * Working directory of the program. If not supplied, the program will run in the current directory.
     * 
     */
    @Import(name="workingDir")
    private @Nullable String workingDir;

    /**
     * @return Working directory of the program. If not supplied, the program will run in the current directory.
     * 
     */
    public Optional<String> workingDir() {
        return Optional.ofNullable(this.workingDir);
    }

    private GetExternalPlainArgs() {}

    private GetExternalPlainArgs(GetExternalPlainArgs $) {
        this.programs = $.programs;
        this.query = $.query;
        this.workingDir = $.workingDir;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetExternalPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetExternalPlainArgs $;

        public Builder() {
            $ = new GetExternalPlainArgs();
        }

        public Builder(GetExternalPlainArgs defaults) {
            $ = new GetExternalPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder programs(List<String> programs) {
            $.programs = programs;
            return this;
        }

        public Builder programs(String... programs) {
            return programs(List.of(programs));
        }

        /**
         * @param query A map of string values to pass to the external program as the query arguments. If not supplied, the program will receive an empty object as its input.
         * 
         * @return builder
         * 
         */
        public Builder query(@Nullable Map<String,String> query) {
            $.query = query;
            return this;
        }

        /**
         * @param workingDir Working directory of the program. If not supplied, the program will run in the current directory.
         * 
         * @return builder
         * 
         */
        public Builder workingDir(@Nullable String workingDir) {
            $.workingDir = workingDir;
            return this;
        }

        public GetExternalPlainArgs build() {
            if ($.programs == null) {
                throw new MissingRequiredPropertyException("GetExternalPlainArgs", "programs");
            }
            return $;
        }
    }

}
